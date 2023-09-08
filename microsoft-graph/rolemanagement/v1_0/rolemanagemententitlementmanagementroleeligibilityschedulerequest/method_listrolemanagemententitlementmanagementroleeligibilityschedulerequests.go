package rolemanagemententitlementmanagementroleeligibilityschedulerequest

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

type ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

type ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequests ...
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequests(ctx context.Context) (result ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests",
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

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsComplete(ctx context.Context) (ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequests(ctx)
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

	result = ListRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult{
		Items: items,
	}
	return
}
