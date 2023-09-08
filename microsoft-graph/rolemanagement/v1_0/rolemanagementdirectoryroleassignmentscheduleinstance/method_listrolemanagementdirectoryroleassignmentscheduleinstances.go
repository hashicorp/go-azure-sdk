package rolemanagementdirectoryroleassignmentscheduleinstance

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

type ListRoleManagementDirectoryRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse
}

type ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteResult struct {
	Items []models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse
}

// ListRoleManagementDirectoryRoleAssignmentScheduleInstances ...
func (c RoleManagementDirectoryRoleAssignmentScheduleInstanceClient) ListRoleManagementDirectoryRoleAssignmentScheduleInstances(ctx context.Context) (result ListRoleManagementDirectoryRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleAssignmentScheduleInstances",
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
		Values *[]models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementDirectoryRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleAssignmentScheduleInstanceClient) ListRoleManagementDirectoryRoleAssignmentScheduleInstancesComplete(ctx context.Context) (ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentScheduleInstanceCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleAssignmentScheduleInstanceClient) ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentScheduleInstanceCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleAssignmentScheduleInstances(ctx)
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

	result = ListRoleManagementDirectoryRoleAssignmentScheduleInstancesCompleteResult{
		Items: items,
	}
	return
}
