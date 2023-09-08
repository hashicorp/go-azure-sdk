package rolemanagemententitlementmanagementroleeligibilityscheduleinstance

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

type ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse
}

type ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse
}

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstances ...
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstances(ctx context.Context) (result ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances",
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
		Values *[]models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesComplete(ctx context.Context) (ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, models.UnifiedRoleEligibilityScheduleInstanceCollectionResponseOperationPredicate{})
}

// ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient) ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleEligibilityScheduleInstanceCollectionResponseOperationPredicate) (result ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse, 0)

	resp, err := c.ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstances(ctx)
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

	result = ListRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult{
		Items: items,
	}
	return
}
