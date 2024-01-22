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

type ListByDepartmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ListByDepartmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

// ListByDepartment ...
func (c BillingRoleAssignmentsClient) ListByDepartment(ctx context.Context, id DepartmentId) (result ListByDepartmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/billingRoleAssignments", id.ID()),
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

// ListByDepartmentComplete retrieves all the results into a single object
func (c BillingRoleAssignmentsClient) ListByDepartmentComplete(ctx context.Context, id DepartmentId) (ListByDepartmentCompleteResult, error) {
	return c.ListByDepartmentCompleteMatchingPredicate(ctx, id, BillingRoleAssignmentOperationPredicate{})
}

// ListByDepartmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentsClient) ListByDepartmentCompleteMatchingPredicate(ctx context.Context, id DepartmentId, predicate BillingRoleAssignmentOperationPredicate) (result ListByDepartmentCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.ListByDepartment(ctx, id)
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

	result = ListByDepartmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
