package profilepublication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProfilePublicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemPublication
}

type ListProfilePublicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemPublication
}

type ListProfilePublicationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfilePublicationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfilePublications ...
func (c ProfilePublicationClient) ListProfilePublications(ctx context.Context, id UserId) (result ListProfilePublicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfilePublicationsCustomPager{},
		Path:       fmt.Sprintf("%s/profile/publications", id.ID()),
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
		Values *[]beta.ItemPublication `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfilePublicationsComplete retrieves all the results into a single object
func (c ProfilePublicationClient) ListProfilePublicationsComplete(ctx context.Context, id UserId) (ListProfilePublicationsCompleteResult, error) {
	return c.ListProfilePublicationsCompleteMatchingPredicate(ctx, id, ItemPublicationOperationPredicate{})
}

// ListProfilePublicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfilePublicationClient) ListProfilePublicationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ItemPublicationOperationPredicate) (result ListProfilePublicationsCompleteResult, err error) {
	items := make([]beta.ItemPublication, 0)

	resp, err := c.ListProfilePublications(ctx, id)
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

	result = ListProfilePublicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
