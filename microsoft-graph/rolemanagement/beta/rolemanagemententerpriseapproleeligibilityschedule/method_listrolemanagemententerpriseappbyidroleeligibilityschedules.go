package rolemanagemententerpriseapproleeligibilityschedule

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

type ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UnifiedRoleEligibilityScheduleCollectionResponse
}

type ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteResult struct {
	Items []models.UnifiedRoleEligibilityScheduleCollectionResponse
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedules ...
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleClient) ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedules(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/roleEligibilitySchedules", id.ID()),
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
		Values *[]models.UnifiedRoleEligibilityScheduleCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesComplete retrieves all the results into a single object
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleClient) ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteResult, error) {
	return c.ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteMatchingPredicate(ctx, id, models.UnifiedRoleEligibilityScheduleCollectionResponseOperationPredicate{})
}

// ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementEnterpriseAppRoleEligibilityScheduleClient) ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate models.UnifiedRoleEligibilityScheduleCollectionResponseOperationPredicate) (result ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteResult, err error) {
	items := make([]models.UnifiedRoleEligibilityScheduleCollectionResponse, 0)

	resp, err := c.ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedules(ctx, id)
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

	result = ListRoleManagementEnterpriseAppByIdRoleEligibilitySchedulesCompleteResult{
		Items: items,
	}
	return
}
