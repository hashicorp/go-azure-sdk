package entitlementmanagementroleeligibilityschedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementRoleEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilitySchedule
}

type ListEntitlementManagementRoleEligibilitySchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilitySchedule
}

type ListEntitlementManagementRoleEligibilitySchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleEligibilitySchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleEligibilitySchedules ...
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedules(ctx context.Context) (result ListEntitlementManagementRoleEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementRoleEligibilitySchedulesCustomPager{},
		Path:       "/roleManagement/entitlementManagement/roleEligibilitySchedules",
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
		Values *[]stable.UnifiedRoleEligibilitySchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleEligibilitySchedulesComplete retrieves all the results into a single object
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedulesComplete(ctx context.Context) (ListEntitlementManagementRoleEligibilitySchedulesCompleteResult, error) {
	return c.ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate(ctx, UnifiedRoleEligibilityScheduleOperationPredicate{})
}

// ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleEligibilityScheduleOperationPredicate) (result ListEntitlementManagementRoleEligibilitySchedulesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilitySchedule, 0)

	resp, err := c.ListEntitlementManagementRoleEligibilitySchedules(ctx)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListEntitlementManagementRoleEligibilitySchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
