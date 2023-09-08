package rolemanagementdirectoryroleeligibilityscheduleinstance

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

type ListRoleManagementDirectoryRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse
}

type ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse
}

// ListRoleManagementDirectoryRoleEligibilityScheduleInstances ...
func (c RoleManagementDirectoryRoleEligibilityScheduleInstanceClient) ListRoleManagementDirectoryRoleEligibilityScheduleInstances(ctx context.Context) (result ListRoleManagementDirectoryRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleEligibilityScheduleInstances",
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

// ListRoleManagementDirectoryRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleEligibilityScheduleInstanceClient) ListRoleManagementDirectoryRoleEligibilityScheduleInstancesComplete(ctx context.Context) (ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, models.UnifiedRoleEligibilityScheduleInstanceCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleEligibilityScheduleInstanceClient) ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleEligibilityScheduleInstanceCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleInstanceCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleEligibilityScheduleInstances(ctx)
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

	result = ListRoleManagementDirectoryRoleEligibilityScheduleInstancesCompleteResult{
		Items: items,
	}
	return
}
