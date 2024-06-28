package topology

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHomeRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TopologyResource
}

type ListByHomeRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TopologyResource
}

type ListByHomeRegionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByHomeRegionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByHomeRegion ...
func (c TopologyClient) ListByHomeRegion(ctx context.Context, id LocationId) (result ListByHomeRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByHomeRegionCustomPager{},
		Path:       fmt.Sprintf("%s/topologies", id.ID()),
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
		Values *[]TopologyResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHomeRegionComplete retrieves all the results into a single object
func (c TopologyClient) ListByHomeRegionComplete(ctx context.Context, id LocationId) (ListByHomeRegionCompleteResult, error) {
	return c.ListByHomeRegionCompleteMatchingPredicate(ctx, id, TopologyResourceOperationPredicate{})
}

// ListByHomeRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TopologyClient) ListByHomeRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate TopologyResourceOperationPredicate) (result ListByHomeRegionCompleteResult, err error) {
	items := make([]TopologyResource, 0)

	resp, err := c.ListByHomeRegion(ctx, id)
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

	result = ListByHomeRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
