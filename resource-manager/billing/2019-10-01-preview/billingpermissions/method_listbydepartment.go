package billingpermissions

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
	Model        *[]BillingPermissionsProperties
}

type ListByDepartmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingPermissionsProperties
}

// ListByDepartment ...
func (c BillingPermissionsClient) ListByDepartment(ctx context.Context, id DepartmentId) (result ListByDepartmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/billingPermissions", id.ID()),
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
		Values *[]BillingPermissionsProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDepartmentComplete retrieves all the results into a single object
func (c BillingPermissionsClient) ListByDepartmentComplete(ctx context.Context, id DepartmentId) (ListByDepartmentCompleteResult, error) {
	return c.ListByDepartmentCompleteMatchingPredicate(ctx, id, BillingPermissionsPropertiesOperationPredicate{})
}

// ListByDepartmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingPermissionsClient) ListByDepartmentCompleteMatchingPredicate(ctx context.Context, id DepartmentId, predicate BillingPermissionsPropertiesOperationPredicate) (result ListByDepartmentCompleteResult, err error) {
	items := make([]BillingPermissionsProperties, 0)

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
