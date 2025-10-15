package dataflowgraph

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataflowProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataflowGraphResource
}

type ListByDataflowProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataflowGraphResource
}

type ListByDataflowProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDataflowProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDataflowProfile ...
func (c DataflowGraphClient) ListByDataflowProfile(ctx context.Context, id DataflowProfileId) (result ListByDataflowProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByDataflowProfileCustomPager{},
		Path:       fmt.Sprintf("%s/dataflowGraphs", id.ID()),
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
		Values *[]DataflowGraphResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDataflowProfileComplete retrieves all the results into a single object
func (c DataflowGraphClient) ListByDataflowProfileComplete(ctx context.Context, id DataflowProfileId) (ListByDataflowProfileCompleteResult, error) {
	return c.ListByDataflowProfileCompleteMatchingPredicate(ctx, id, DataflowGraphResourceOperationPredicate{})
}

// ListByDataflowProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataflowGraphClient) ListByDataflowProfileCompleteMatchingPredicate(ctx context.Context, id DataflowProfileId, predicate DataflowGraphResourceOperationPredicate) (result ListByDataflowProfileCompleteResult, err error) {
	items := make([]DataflowGraphResource, 0)

	resp, err := c.ListByDataflowProfile(ctx, id)
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

	result = ListByDataflowProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
