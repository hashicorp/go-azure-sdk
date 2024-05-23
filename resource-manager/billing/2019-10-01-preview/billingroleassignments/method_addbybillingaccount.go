package billingroleassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type AddByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

// AddByBillingAccount ...
func (c BillingRoleAssignmentsClient) AddByBillingAccount(ctx context.Context, id BillingAccountId, input BillingRoleAssignmentPayload) (result AddByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/createBillingRoleAssignment", id.ID()),
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
		Values *[]BillingRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AddByBillingAccountComplete retrieves all the results into a single object
func (c BillingRoleAssignmentsClient) AddByBillingAccountComplete(ctx context.Context, id BillingAccountId, input BillingRoleAssignmentPayload) (AddByBillingAccountCompleteResult, error) {
	return c.AddByBillingAccountCompleteMatchingPredicate(ctx, id, input, BillingRoleAssignmentOperationPredicate{})
}

// AddByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentsClient) AddByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, input BillingRoleAssignmentPayload, predicate BillingRoleAssignmentOperationPredicate) (result AddByBillingAccountCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.AddByBillingAccount(ctx, id, input)
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

	result = AddByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
