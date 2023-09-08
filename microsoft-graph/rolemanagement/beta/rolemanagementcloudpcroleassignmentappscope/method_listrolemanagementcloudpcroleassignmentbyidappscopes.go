package rolemanagementcloudpcroleassignmentappscope

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

type ListRoleManagementCloudPCRoleAssignmentByIdAppScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppScopeCollectionResponse
}

type ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteResult struct {
	Items []models.AppScopeCollectionResponse
}

// ListRoleManagementCloudPCRoleAssignmentByIdAppScopes ...
func (c RoleManagementCloudPCRoleAssignmentAppScopeClient) ListRoleManagementCloudPCRoleAssignmentByIdAppScopes(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (result ListRoleManagementCloudPCRoleAssignmentByIdAppScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appScopes", id.ID()),
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
		Values *[]models.AppScopeCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementCloudPCRoleAssignmentByIdAppScopesComplete retrieves all the results into a single object
func (c RoleManagementCloudPCRoleAssignmentAppScopeClient) ListRoleManagementCloudPCRoleAssignmentByIdAppScopesComplete(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteResult, error) {
	return c.ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteMatchingPredicate(ctx, id, models.AppScopeCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCRoleAssignmentAppScopeClient) ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteMatchingPredicate(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId, predicate models.AppScopeCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteResult, err error) {
	items := make([]models.AppScopeCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCRoleAssignmentByIdAppScopes(ctx, id)
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

	result = ListRoleManagementCloudPCRoleAssignmentByIdAppScopesCompleteResult{
		Items: items,
	}
	return
}
