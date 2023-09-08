package rolemanagementdirectoryroleeligibilityschedulerequest

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

type ListRoleManagementDirectoryRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

type ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleRequestCollectionResponse
}

// ListRoleManagementDirectoryRoleEligibilityScheduleRequests ...
func (c RoleManagementDirectoryRoleEligibilityScheduleRequestClient) ListRoleManagementDirectoryRoleEligibilityScheduleRequests(ctx context.Context) (result ListRoleManagementDirectoryRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleEligibilityScheduleRequests",
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

// ListRoleManagementDirectoryRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c RoleManagementDirectoryRoleEligibilityScheduleRequestClient) ListRoleManagementDirectoryRoleEligibilityScheduleRequestsComplete(ctx context.Context) (ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate{})
}

// ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementDirectoryRoleEligibilityScheduleRequestClient) ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate models.UnifiedRoleEligibilityScheduleRequestCollectionResponseOperationPredicate) (result ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleRequestCollectionResponse, 0)

	resp, err := c.ListRoleManagementDirectoryRoleEligibilityScheduleRequests(ctx)
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

	result = ListRoleManagementDirectoryRoleEligibilityScheduleRequestsCompleteResult{
		Items: items,
	}
	return
}
