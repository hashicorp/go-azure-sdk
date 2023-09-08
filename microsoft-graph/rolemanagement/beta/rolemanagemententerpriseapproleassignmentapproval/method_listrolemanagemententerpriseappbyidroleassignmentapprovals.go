package rolemanagemententerpriseapproleassignmentapproval

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

type ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteResult struct {
	Items []models.ApprovalCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovals ...
func (c RoleManagementEnterpriseAppRoleAssignmentApprovalClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovals(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roleAssignmentApprovals", id.ID()),
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

// ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppRoleAssignmentApprovalClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, id, models.ApprovalCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppRoleAssignmentApprovalClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.ApprovalCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]models.ApprovalCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovals(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdRoleAssignmentApprovalsCompleteResult{
		Items: items,
	}
	return
}
