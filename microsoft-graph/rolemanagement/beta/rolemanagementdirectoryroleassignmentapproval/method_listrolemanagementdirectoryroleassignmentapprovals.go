package rolemanagementdirectoryroleassignmentapproval

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

type ListRoleManagementDirectoryRoleAssignmentApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ApprovalCollectionResponse
}

type ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteResult struct {
	Items []models.ApprovalCollectionResponse
}

// ListRoleManagementDirectoryRoleAssignmentApprovals ...
func (c RoleManagementDirectoryRoleAssignmentApprovalClient) ListRoleManagementDirectoryRoleAssignmentApprovals(ctx context.Context) (result ListRoleManagementDirectoryRoleAssignmentApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleAssignmentApprovals",
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

// ListRoleManagementDirectoryRoleAssignmentApprovalsComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleAssignmentApprovalClient) ListRoleManagementDirectoryRoleAssignmentApprovalsComplete(ctx context.Context) (ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate(ctx, models.ApprovalCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleAssignmentApprovalClient) ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteMatchingPredicate(ctx context.Context, predicate models.ApprovalCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteResult, err error) {
	items := make([]models.ApprovalCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleAssignmentApprovals(ctx)
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

	result = ListRoleManagementDirectoryRoleAssignmentApprovalsCompleteResult{
		Items: items,
	}
	return
}
