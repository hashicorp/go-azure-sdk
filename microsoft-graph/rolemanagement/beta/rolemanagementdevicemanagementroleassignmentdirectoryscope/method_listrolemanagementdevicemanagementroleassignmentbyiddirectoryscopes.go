package rolemanagementdevicemanagementroleassignmentdirectoryscope

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

type ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopes ...
func (c RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient) ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopes(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (result ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/directoryScopes", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesComplete retrieves all the results into a single object
func (c RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient) ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesComplete(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteResult, error) {
	return c.ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient) ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopes(ctx, id)
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

	result = ListRoleManagementDeviceManagementRoleAssignmentByIdDirectoryScopesCompleteResult{
		Items: items,
	}
	return
}
