package virtualnetworkrules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualNetworkRule
}

type ListByAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualNetworkRule
}

// ListByAccount ...
func (c VirtualNetworkRulesClient) ListByAccount(ctx context.Context, id AccountId) (result ListByAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/virtualNetworkRules", id.ID()),
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
		Values *[]VirtualNetworkRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAccountComplete retrieves all the results into a single object
func (c VirtualNetworkRulesClient) ListByAccountComplete(ctx context.Context, id AccountId) (ListByAccountCompleteResult, error) {
	return c.ListByAccountCompleteMatchingPredicate(ctx, id, VirtualNetworkRuleOperationPredicate{})
}

// ListByAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualNetworkRulesClient) ListByAccountCompleteMatchingPredicate(ctx context.Context, id AccountId, predicate VirtualNetworkRuleOperationPredicate) (result ListByAccountCompleteResult, err error) {
	items := make([]VirtualNetworkRule, 0)

	resp, err := c.ListByAccount(ctx, id)
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

	result = ListByAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
