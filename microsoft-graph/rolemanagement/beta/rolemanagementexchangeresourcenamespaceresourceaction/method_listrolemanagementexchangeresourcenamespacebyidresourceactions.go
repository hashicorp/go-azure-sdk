package rolemanagementexchangeresourcenamespaceresourceaction

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

type ListRoleManagementExchangeResourceNamespaceByIdResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRbacResourceActionCollectionResponse
}

type ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteResult struct {
	Items []models.UnifiedRbacResourceActionCollectionResponse
}

// ListRoleManagementExchangeResourceNamespaceByIdResourceActions ...
func (c RoleManagementExchangeResourceNamespaceResourceActionClient) ListRoleManagementExchangeResourceNamespaceByIdResourceActions(ctx context.Context, id RoleManagementExchangeResourceNamespaceId) (result ListRoleManagementExchangeResourceNamespaceByIdResourceActionsOperationResponse, err error) {
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

// ListRoleManagementExchangeResourceNamespaceByIdResourceActionsComplete retrieves all the results into a single object
func (c RoleManagementExchangeResourceNamespaceResourceActionClient) ListRoleManagementExchangeResourceNamespaceByIdResourceActionsComplete(ctx context.Context, id RoleManagementExchangeResourceNamespaceId) (ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteResult, error) {
	return c.ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx, id, models.UnifiedRbacResourceActionCollectionResponseOperationPredicate{})
}

// ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementExchangeResourceNamespaceResourceActionClient) ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementExchangeResourceNamespaceId, predicate models.UnifiedRbacResourceActionCollectionResponseOperationPredicate) (result ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteResult, err error) {
	items := make([]models.UnifiedRbacResourceActionCollectionResponse, 0)

	resp, err := c.ListRoleManagementExchangeResourceNamespaceByIdResourceActions(ctx, id)
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

	result = ListRoleManagementExchangeResourceNamespaceByIdResourceActionsCompleteResult{
		Items: items,
	}
	return
}
