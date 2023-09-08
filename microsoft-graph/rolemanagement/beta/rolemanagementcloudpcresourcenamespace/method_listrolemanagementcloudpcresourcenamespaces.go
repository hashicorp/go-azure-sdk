package rolemanagementcloudpcresourcenamespace

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

type ListRoleManagementCloudPCResourceNamespacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRbacResourceNamespaceCollectionResponse
}

type ListRoleManagementCloudPCResourceNamespacesCompleteResult struct {
	Items []models.UnifiedRbacResourceNamespaceCollectionResponse
}

// ListRoleManagementCloudPCResourceNamespaces ...
func (c RoleManagementCloudPCResourceNamespaceClient) ListRoleManagementCloudPCResourceNamespaces(ctx context.Context) (result ListRoleManagementCloudPCResourceNamespacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/cloudPC/resourceNamespaces",
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
		Values *[]models.UnifiedRbacResourceNamespaceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementCloudPCResourceNamespacesComplete retrieves all the results into a single object
func (c RoleManagementCloudPCResourceNamespaceClient) ListRoleManagementCloudPCResourceNamespacesComplete(ctx context.Context) (ListRoleManagementCloudPCResourceNamespacesCompleteResult, error) {
	return c.ListRoleManagementCloudPCResourceNamespacesCompleteMatchingPredicate(ctx, models.UnifiedRbacResourceNamespaceCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCResourceNamespacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCResourceNamespaceClient) ListRoleManagementCloudPCResourceNamespacesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRbacResourceNamespaceCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCResourceNamespacesCompleteResult, err error) {
	items := make([]models.UnifiedRbacResourceNamespaceCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCResourceNamespaces(ctx)
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

	result = ListRoleManagementCloudPCResourceNamespacesCompleteResult{
		Items: items,
	}
	return
}
