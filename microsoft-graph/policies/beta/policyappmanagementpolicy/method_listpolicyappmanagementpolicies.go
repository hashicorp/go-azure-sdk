package policyappmanagementpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppManagementPolicyCollectionResponse
}

type ListPolicyAppManagementPoliciesCompleteResult struct {
	Items []models.AppManagementPolicyCollectionResponse
}

// ListPolicyAppManagementPolicies ...
func (c PolicyAppManagementPolicyClient) ListPolicyAppManagementPolicies(ctx context.Context) (result ListPolicyAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/appManagementPolicies",
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
		Values *[]models.AppManagementPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyAppManagementPoliciesComplete retrieves all the results into a single object
func (c PolicyAppManagementPolicyClient) ListPolicyAppManagementPoliciesComplete(ctx context.Context) (ListPolicyAppManagementPoliciesCompleteResult, error) {
	return c.ListPolicyAppManagementPoliciesCompleteMatchingPredicate(ctx, models.AppManagementPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyAppManagementPolicyClient) ListPolicyAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.AppManagementPolicyCollectionResponseOperationPredicate) (result ListPolicyAppManagementPoliciesCompleteResult, err error) {
	items := make([]models.AppManagementPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyAppManagementPolicies(ctx)
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

	result = ListPolicyAppManagementPoliciesCompleteResult{
		Items: items,
	}
	return
}
