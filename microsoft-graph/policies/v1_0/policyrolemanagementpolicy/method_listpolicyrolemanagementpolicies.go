package policyrolemanagementpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyRoleManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleManagementPolicyCollectionResponse
}

type ListPolicyRoleManagementPoliciesCompleteResult struct {
	Items []models.UnifiedRoleManagementPolicyCollectionResponse
}

// ListPolicyRoleManagementPolicies ...
func (c PolicyRoleManagementPolicyClient) ListPolicyRoleManagementPolicies(ctx context.Context) (result ListPolicyRoleManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/roleManagementPolicies",
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
		Values *[]models.UnifiedRoleManagementPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyRoleManagementPoliciesComplete retrieves all the results into a single object
func (c PolicyRoleManagementPolicyClient) ListPolicyRoleManagementPoliciesComplete(ctx context.Context) (ListPolicyRoleManagementPoliciesCompleteResult, error) {
	return c.ListPolicyRoleManagementPoliciesCompleteMatchingPredicate(ctx, models.UnifiedRoleManagementPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyRoleManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyRoleManagementPolicyClient) ListPolicyRoleManagementPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleManagementPolicyCollectionResponseOperationPredicate) (result ListPolicyRoleManagementPoliciesCompleteResult, err error) {
	items := make([]models.UnifiedRoleManagementPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyRoleManagementPolicies(ctx)
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

	result = ListPolicyRoleManagementPoliciesCompleteResult{
		Items: items,
	}
	return
}
