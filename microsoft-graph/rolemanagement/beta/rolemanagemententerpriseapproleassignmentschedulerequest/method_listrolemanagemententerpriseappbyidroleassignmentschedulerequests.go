package rolemanagemententerpriseapproleassignmentschedulerequest

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

type ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleAssignmentScheduleRequestCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteResult struct {
	Items []models.UnifiedRoleAssignmentScheduleRequestCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequests ...
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequests(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roleAssignmentScheduleRequests", id.ID()),
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
		Values *[]models.UnifiedRoleAssignmentScheduleRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleAssignmentScheduleRequestCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppRoleAssignmentScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.UnifiedRoleAssignmentScheduleRequestCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]models.UnifiedRoleAssignmentScheduleRequestCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequests(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdRoleAssignmentScheduleRequestsCompleteResult{
		Items: items,
	}
	return
}
