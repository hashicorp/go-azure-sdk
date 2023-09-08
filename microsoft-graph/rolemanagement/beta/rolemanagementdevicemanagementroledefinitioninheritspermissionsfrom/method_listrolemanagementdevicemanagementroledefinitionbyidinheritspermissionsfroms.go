package rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom

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

type ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleDefinitionCollectionResponse
}

type ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteResult struct {
	Items []models.UnifiedRoleDefinitionCollectionResponse
}

// ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFroms ...
func (c RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFroms(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId) (result ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsOperationResponse, err error) {
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

// ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId) (ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, error) {
	return c.ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleDefinitionCollectionResponseOperationPredicate{})
}

// ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient) ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleDefinitionId, predicate models.UnifiedRoleDefinitionCollectionResponseOperationPredicate) (result ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]models.UnifiedRoleDefinitionCollectionResponse, 0)

	resp, err := c.ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFroms(ctx, id)
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

	result = ListRoleManagementDeviceManagementRoleDefinitionByIdInheritsPermissionsFromsCompleteResult{
		Items: items,
	}
	return
}
