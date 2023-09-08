package rolemanagementdirectoryroleassignmentschedule

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

type ListRoleManagementDirectoryRoleAssignmentSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentScheduleCollectionResponse
}

type ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteResult struct {
	Items []models.UnifiedRoleAssignmentScheduleCollectionResponse
}

// ListRoleManagementDirectoryRoleAssignmentSchedules ...
func (c RoleManagementDirectoryRoleAssignmentScheduleClient) ListRoleManagementDirectoryRoleAssignmentSchedules(ctx context.Context) (result ListRoleManagementDirectoryRoleAssignmentSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleAssignmentSchedules",
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
		Values *[]models.UnifiedRoleAssignmentScheduleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementDirectoryRoleAssignmentSchedulesComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleAssignmentScheduleClient) ListRoleManagementDirectoryRoleAssignmentSchedulesComplete(ctx context.Context) (ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate(ctx, models.UnifiedRoleAssignmentScheduleCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleAssignmentScheduleClient) ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleAssignmentScheduleCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentScheduleCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleAssignmentSchedules(ctx)
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

	result = ListRoleManagementDirectoryRoleAssignmentSchedulesCompleteResult{
		Items: items,
	}
	return
}
