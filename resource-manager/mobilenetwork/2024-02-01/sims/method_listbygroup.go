package sims

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Sim
}

type ListByGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Sim
}

// ListByGroup ...
func (c SIMsClient) ListByGroup(ctx context.Context, id SimGroupId) (result ListByGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sims", id.ID()),
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
		Values *[]Sim `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByGroupComplete retrieves all the results into a single object
func (c SIMsClient) ListByGroupComplete(ctx context.Context, id SimGroupId) (ListByGroupCompleteResult, error) {
	return c.ListByGroupCompleteMatchingPredicate(ctx, id, SimOperationPredicate{})
}

// ListByGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SIMsClient) ListByGroupCompleteMatchingPredicate(ctx context.Context, id SimGroupId, predicate SimOperationPredicate) (result ListByGroupCompleteResult, err error) {
	items := make([]Sim, 0)

	resp, err := c.ListByGroup(ctx, id)
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

	result = ListByGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
