package rolemanagementdirectoryresourcenamespaceresourceaction

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

type ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRbacResourceActionCollectionResponse
}

type ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteResult struct {
	Items []models.UnifiedRbacResourceActionCollectionResponse
}

// ListRoleManagementDirectoryResourceNamespaceByIdResourceActions ...
func (c RoleManagementDirectoryResourceNamespaceResourceActionClient) ListRoleManagementDirectoryResourceNamespaceByIdResourceActions(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId) (result ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsOperationResponse, err error) {
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

// ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsComplete retrieves all the results into a single object
func (c RoleManagementDirectoryResourceNamespaceResourceActionClient) ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsComplete(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId) (ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteResult, error) {
	return c.ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx, id, models.UnifiedRbacResourceActionCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryResourceNamespaceResourceActionClient) ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId, predicate models.UnifiedRbacResourceActionCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteResult, err error) {
	items := make([]models.UnifiedRbacResourceActionCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryResourceNamespaceByIdResourceActions(ctx, id)
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

	result = ListRoleManagementDirectoryResourceNamespaceByIdResourceActionsCompleteResult{
		Items: items,
	}
	return
}
