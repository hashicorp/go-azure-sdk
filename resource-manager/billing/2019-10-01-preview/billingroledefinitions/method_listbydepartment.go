package billingroledefinitions

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
	Model        *[]BillingRoleDefinition
}

type ListByDepartmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleDefinition
}

// ListByDepartment ...
func (c BillingRoleDefinitionsClient) ListByDepartment(ctx context.Context, id DepartmentId) (result ListByDepartmentOperationResponse, err error) {
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

// ListByDepartmentComplete retrieves all the results into a single object
func (c BillingRoleDefinitionsClient) ListByDepartmentComplete(ctx context.Context, id DepartmentId) (ListByDepartmentCompleteResult, error) {
	return c.ListByDepartmentCompleteMatchingPredicate(ctx, id, BillingRoleDefinitionOperationPredicate{})
}

// ListByDepartmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleDefinitionsClient) ListByDepartmentCompleteMatchingPredicate(ctx context.Context, id DepartmentId, predicate BillingRoleDefinitionOperationPredicate) (result ListByDepartmentCompleteResult, err error) {
	items := make([]BillingRoleDefinition, 0)

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
