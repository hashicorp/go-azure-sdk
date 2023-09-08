package rolemanagementcloudpcroleassignment

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

type ListRoleManagementCloudPCRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentMultipleCollectionResponse
}

type ListRoleManagementCloudPCRoleAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentMultipleCollectionResponse
}

// ListRoleManagementCloudPCRoleAssignments ...
func (c RoleManagementCloudPCRoleAssignmentClient) ListRoleManagementCloudPCRoleAssignments(ctx context.Context) (result ListRoleManagementCloudPCRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/cloudPC/roleAssignments",
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
		Values *[]models.UnifiedRoleAssignmentMultipleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementCloudPCRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementCloudPCRoleAssignmentClient) ListRoleManagementCloudPCRoleAssignmentsComplete(ctx context.Context) (ListRoleManagementCloudPCRoleAssignmentsCompleteResult, error) {
	return c.ListRoleManagementCloudPCRoleAssignmentsCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentMultipleCollectionResponseOperationPredicate{})
}

// ListRoleManagementCloudPCRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementCloudPCRoleAssignmentClient) ListRoleManagementCloudPCRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentMultipleCollectionResponseOperationPredicate) (result ListRoleManagementCloudPCRoleAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentMultipleCollectionResponse, 0)

	resp, err := c.ListRoleManagementCloudPCRoleAssignments(ctx)
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

	result = ListRoleManagementCloudPCRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}
