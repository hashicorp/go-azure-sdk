package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeNodesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EdgeNode
}

type EdgeNodesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EdgeNode
}

type EdgeNodesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EdgeNodesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EdgeNodesList ...
func (c OpenapisClient) EdgeNodesList(ctx context.Context) (result EdgeNodesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &EdgeNodesListCustomPager{},
		Path:       "/providers/Microsoft.Cdn/edgenodes",
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
		Values *[]EdgeNode `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// EdgeNodesListComplete retrieves all the results into a single object
func (c OpenapisClient) EdgeNodesListComplete(ctx context.Context) (EdgeNodesListCompleteResult, error) {
	return c.EdgeNodesListCompleteMatchingPredicate(ctx, EdgeNodeOperationPredicate{})
}

// EdgeNodesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) EdgeNodesListCompleteMatchingPredicate(ctx context.Context, predicate EdgeNodeOperationPredicate) (result EdgeNodesListCompleteResult, err error) {
	items := make([]EdgeNode, 0)

	resp, err := c.EdgeNodesList(ctx)
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

	result = EdgeNodesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
