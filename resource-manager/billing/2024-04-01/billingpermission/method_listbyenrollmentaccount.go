package billingpermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByEnrollmentAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingPermission
}

type ListByEnrollmentAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingPermission
}

// ListByEnrollmentAccount ...
func (c BillingPermissionClient) ListByEnrollmentAccount(ctx context.Context, id EnrollmentAccountId) (result ListByEnrollmentAccountOperationResponse, err error) {
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
		Values *[]BillingPermission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByEnrollmentAccountComplete retrieves all the results into a single object
func (c BillingPermissionClient) ListByEnrollmentAccountComplete(ctx context.Context, id EnrollmentAccountId) (ListByEnrollmentAccountCompleteResult, error) {
	return c.ListByEnrollmentAccountCompleteMatchingPredicate(ctx, id, BillingPermissionOperationPredicate{})
}

// ListByEnrollmentAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingPermissionClient) ListByEnrollmentAccountCompleteMatchingPredicate(ctx context.Context, id EnrollmentAccountId, predicate BillingPermissionOperationPredicate) (result ListByEnrollmentAccountCompleteResult, err error) {
	items := make([]BillingPermission, 0)

	resp, err := c.ListByEnrollmentAccount(ctx, id)
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

	result = ListByEnrollmentAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
