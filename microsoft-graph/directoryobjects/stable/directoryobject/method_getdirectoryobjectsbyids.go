package directoryobject

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

type GetDirectoryObjectsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type GetDirectoryObjectsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type GetDirectoryObjectsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDirectoryObjectsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDirectoryObjectsByIds ...
func (c DirectoryObjectClient) GetDirectoryObjectsByIds(ctx context.Context, input GetDirectoryObjectsByIdsRequest) (result GetDirectoryObjectsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetDirectoryObjectsByIdsCustomPager{},
		Path:       "/directoryObjects/getByIds",
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
		Values *[]stable.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDirectoryObjectsByIdsComplete retrieves all the results into a single object
func (c DirectoryObjectClient) GetDirectoryObjectsByIdsComplete(ctx context.Context, input GetDirectoryObjectsByIdsRequest) (GetDirectoryObjectsByIdsCompleteResult, error) {
	return c.GetDirectoryObjectsByIdsCompleteMatchingPredicate(ctx, input, DirectoryObjectOperationPredicate{})
}

// GetDirectoryObjectsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryObjectClient) GetDirectoryObjectsByIdsCompleteMatchingPredicate(ctx context.Context, input GetDirectoryObjectsByIdsRequest, predicate DirectoryObjectOperationPredicate) (result GetDirectoryObjectsByIdsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.GetDirectoryObjectsByIds(ctx, input)
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

	result = GetDirectoryObjectsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
