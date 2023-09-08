package rolemanagementdevicemanagementroleassignmentprincipal

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

type ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipals ...
func (c RoleManagementDeviceManagementRoleAssignmentPrincipalClient) ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipals(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (result ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/principals", id.ID()),
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

// ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsComplete retrieves all the results into a single object
func (c RoleManagementDeviceManagementRoleAssignmentPrincipalClient) ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsComplete(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId) (ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteResult, error) {
	return c.ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDeviceManagementRoleAssignmentPrincipalClient) ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDeviceManagementRoleAssignmentId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipals(ctx, id)
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

	result = ListRoleManagementDeviceManagementRoleAssignmentByIdPrincipalsCompleteResult{
		Items: items,
	}
	return
}
