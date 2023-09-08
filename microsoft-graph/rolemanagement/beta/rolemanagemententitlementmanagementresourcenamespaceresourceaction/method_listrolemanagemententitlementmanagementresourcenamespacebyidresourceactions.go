package rolemanagemententitlementmanagementresourcenamespaceresourceaction

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

type ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRbacResourceActionCollectionResponse
}

type ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteResult struct {
	Items []models.UnifiedRbacResourceActionCollectionResponse
}

// ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActions ...
func (c RoleManagementEntitlementManagementResourceNamespaceResourceActionClient) ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActions(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId) (result ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsOperationResponse, err error) {
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

// ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementResourceNamespaceResourceActionClient) ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsComplete(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId) (ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx, id, models.UnifiedRbacResourceActionCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementResourceNamespaceResourceActionClient) ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId, predicate models.UnifiedRbacResourceActionCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteResult, err error) {
	items := make([]models.UnifiedRbacResourceActionCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActions(ctx, id)
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

	result = ListRoleManagementEntitlementManagementResourceNamespaceByIdResourceActionsCompleteResult{
		Items: items,
	}
	return
}
