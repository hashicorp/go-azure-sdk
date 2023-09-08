package rolemanagemententerpriseapproleeligibilityschedulerequest

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

type ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequests ...
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequests(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roleEligibilityScheduleRequests", id.ID()),
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
		Values *[]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, id, models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient) ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequests(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdRoleEligibilityScheduleRequestsCompleteResult{
		Items: items,
	}
	return
}
