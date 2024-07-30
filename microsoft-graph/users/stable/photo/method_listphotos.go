package photo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPhotosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ProfilePhoto
}

type ListPhotosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ProfilePhoto
}

type ListPhotosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPhotosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPhotos ...
func (c PhotoClient) ListPhotos(ctx context.Context, id UserId) (result ListPhotosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPhotosCustomPager{},
		Path:       fmt.Sprintf("%s/photos", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.ProfilePhoto `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPhotosComplete retrieves all the results into a single object
func (c PhotoClient) ListPhotosComplete(ctx context.Context, id UserId) (ListPhotosCompleteResult, error) {
	return c.ListPhotosCompleteMatchingPredicate(ctx, id, ProfilePhotoOperationPredicate{})
}

// ListPhotosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PhotoClient) ListPhotosCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ProfilePhotoOperationPredicate) (result ListPhotosCompleteResult, err error) {
	items := make([]stable.ProfilePhoto, 0)

	resp, err := c.ListPhotos(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListPhotosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
