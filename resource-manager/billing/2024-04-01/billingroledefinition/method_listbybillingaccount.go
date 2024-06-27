package billingroledefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleDefinition
}

type ListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleDefinition
}

// ListByBillingAccount ...
func (c BillingRoleDefinitionClient) ListByBillingAccount(ctx context.Context, id BillingAccountId) (result ListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/billingRoleDefinitions", id.ID()),
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
		Values *[]BillingRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingAccountComplete retrieves all the results into a single object
func (c BillingRoleDefinitionClient) ListByBillingAccountComplete(ctx context.Context, id BillingAccountId) (ListByBillingAccountCompleteResult, error) {
	return c.ListByBillingAccountCompleteMatchingPredicate(ctx, id, BillingRoleDefinitionOperationPredicate{})
}

// ListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleDefinitionClient) ListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, predicate BillingRoleDefinitionOperationPredicate) (result ListByBillingAccountCompleteResult, err error) {
	items := make([]BillingRoleDefinition, 0)

	resp, err := c.ListByBillingAccount(ctx, id)
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

	result = ListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
