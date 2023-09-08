package rolemanagemententerpriseapproleassignmentscheduleinstance

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

type ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteResult struct {
	Items []models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstances ...
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstances(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roleAssignmentScheduleInstances", id.ID()),
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

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, id, models.UnifiedRoleAssignmentScheduleInstanceCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.UnifiedRoleAssignmentScheduleInstanceCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentScheduleInstanceCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstances(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleInstancesCompleteResult{
		Items: items,
	}
	return
}
