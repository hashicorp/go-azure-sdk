package locationbasedrecommendedactionsessionsresult

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RecommendationAction
}

type ListCompleteResult struct {
	Items []RecommendationAction
}

// List ...
func (c LocationBasedRecommendedActionSessionsResultClient) List(ctx context.Context, id RecommendedActionSessionsOperationResultId) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       id.ID(),
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
		Values *[]RecommendationAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c LocationBasedRecommendedActionSessionsResultClient) ListComplete(ctx context.Context, id RecommendedActionSessionsOperationResultId) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, RecommendationActionOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LocationBasedRecommendedActionSessionsResultClient) ListCompleteMatchingPredicate(ctx context.Context, id RecommendedActionSessionsOperationResultId, predicate RecommendationActionOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]RecommendationAction, 0)

	resp, err := c.List(ctx, id)
	if err != nil {
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

	result = ListCompleteResult{
		Items: items,
	}
	return
}
