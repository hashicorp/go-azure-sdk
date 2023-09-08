package rolemanagementcloudpcroleassignmentprincipal

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

type ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListRoleManagementCloudPCRoleAssignmentByIdPrincipals ...
func (c RoleManagementCloudPCRoleAssignmentPrincipalClient) ListRoleManagementCloudPCRoleAssignmentByIdPrincipals(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (result ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsOperationResponse, err error) {
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

// ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsComplete retrieves all the results into a single object
func (c RoleManagementCloudPCRoleAssignmentPrincipalClient) ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsComplete(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId) (ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteResult, error) {
	return c.ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCRoleAssignmentPrincipalClient) ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteMatchingPredicate(ctx context.Context, id RoleManagementCloudPCRoleAssignmentId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCRoleAssignmentByIdPrincipals(ctx, id)
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

	result = ListRoleManagementCloudPCRoleAssignmentByIdPrincipalsCompleteResult{
		Items: items,
	}
	return
}
