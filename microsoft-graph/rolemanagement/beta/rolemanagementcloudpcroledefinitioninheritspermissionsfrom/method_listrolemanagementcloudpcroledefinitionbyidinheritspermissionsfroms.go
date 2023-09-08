package rolemanagementcloudpcroledefinitioninheritspermissionsfrom

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

type ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFroms ...
func (c RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFroms(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId) (result ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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
		Values *[]models.UnifiedRoleDefinitionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId) (ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, error) {
	return c.ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementCloudPCRoleDefinitionId, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFroms(ctx, id)
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

	result = ListRoleManagementCloudPCRoleDefinitionByIdInheritsPermissionsFromsCompleteResult{
		Items: items,
	}
	return
}
