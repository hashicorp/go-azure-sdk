package rolemanagementdevicemanagementresourcenamespaceresourceaction

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

type ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRbacResourceActionCollectionResponse
}

type ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteResult struct {
	Items []models.UnifiedRbacResourceActionCollectionResponse
}

// ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActions ...
func (c RoleManagementDeviceManagementResourceNamespaceResourceActionClient) ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActions(ctx context.Context, id RoleManagementDeviceManagementResourceNamespaceId) (result ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/resourceActions", id.ID()),
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
		Values *[]models.UnifiedRbacResourceActionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsComplete retrieves all the results into a single object
func (c RoleManagementDeviceManagementResourceNamespaceResourceActionClient) ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsComplete(ctx context.Context, id RoleManagementDeviceManagementResourceNamespaceId) (ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteResult, error) {
	return c.ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx, id, models.UnifiedRbacResourceActionCollectionResponseOperationPredicate{})
}

// ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDeviceManagementResourceNamespaceResourceActionClient) ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementResourceNamespaceId, predicate models.UnifiedRbacResourceActionCollectionResponseOperationPredicate) (result ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteResult, err error) {
	items := make([]models.UnifiedRbacResourceActionCollectionResponse, 0)

	resp, err := c.ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActions(ctx, id)
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

	result = ListRoleManagementDeviceManagementResourceNamespaceByIdResourceActionsCompleteResult{
		Items: items,
	}
	return
}
