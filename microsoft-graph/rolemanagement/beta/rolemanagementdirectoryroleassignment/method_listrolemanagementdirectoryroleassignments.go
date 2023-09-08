package rolemanagementdirectoryroleassignment

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

type ListRoleManagementDirectoryRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentCollectionResponse
}

type ListRoleManagementDirectoryRoleAssignmentsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentCollectionResponse
}

// ListRoleManagementDirectoryRoleAssignments ...
func (c RoleManagementDirectoryRoleAssignmentClient) ListRoleManagementDirectoryRoleAssignments(ctx context.Context) (result ListRoleManagementDirectoryRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleAssignments",
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
		Values *[]models.UnifiedRoleAssignmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementDirectoryRoleAssignmentsComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleAssignmentClient) ListRoleManagementDirectoryRoleAssignmentsComplete(ctx context.Context) (ListRoleManagementDirectoryRoleAssignmentsCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleAssignmentsCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleAssignmentClient) ListRoleManagementDirectoryRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleAssignmentsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleAssignments(ctx)
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

	result = ListRoleManagementDirectoryRoleAssignmentsCompleteResult{
		Items: items,
	}
	return
}
