package rolemanagemententitlementmanagementroleassignmentapproval

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

type ListRoleManagementEntitlementManagementRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalCollectionResponse
}

type ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteResult struct {
	Items []models.ApprovalCollectionResponse
}

// ListRoleManagementEntitlementManagementRoleAssignmentApprovals ...
func (c RoleManagementEntitlementManagementRoleAssignmentApprovalClient) ListRoleManagementEntitlementManagementRoleAssignmentApprovals(ctx context.Context) (result ListRoleManagementEntitlementManagementRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/entitlementManagement/roleAssignmentApprovals",
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
		Values *[]models.ApprovalCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEntitlementManagementRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementRoleAssignmentApprovalClient) ListRoleManagementEntitlementManagementRoleAssignmentApprovalsComplete(ctx context.Context) (ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, models.ApprovalCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementRoleAssignmentApprovalClient) ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate models.ApprovalCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]models.ApprovalCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementRoleAssignmentApprovals(ctx)
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

	result = ListRoleManagementEntitlementManagementRoleAssignmentApprovalsCompleteResult{
		Items: items,
	}
	return
}
