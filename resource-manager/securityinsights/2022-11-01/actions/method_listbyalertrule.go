package actions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAlertRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ActionResponse
}

type ListByAlertRuleCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ActionResponse
}

// ListByAlertRule ...
func (c ActionsClient) ListByAlertRule(ctx context.Context, id AlertRuleId) (result ListByAlertRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/actions", id.ID()),
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
		Values *[]ActionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAlertRuleComplete retrieves all the results into a single object
func (c ActionsClient) ListByAlertRuleComplete(ctx context.Context, id AlertRuleId) (ListByAlertRuleCompleteResult, error) {
	return c.ListByAlertRuleCompleteMatchingPredicate(ctx, id, ActionResponseOperationPredicate{})
}

// ListByAlertRuleCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ActionsClient) ListByAlertRuleCompleteMatchingPredicate(ctx context.Context, id AlertRuleId, predicate ActionResponseOperationPredicate) (result ListByAlertRuleCompleteResult, err error) {
	items := make([]ActionResponse, 0)

	resp, err := c.ListByAlertRule(ctx, id)
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

	result = ListByAlertRuleCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
