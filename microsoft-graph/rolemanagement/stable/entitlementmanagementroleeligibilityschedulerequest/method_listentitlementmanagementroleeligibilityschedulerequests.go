package entitlementmanagementroleeligibilityschedulerequest

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

type ListEntitlementManagementRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleRequest
}

type ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleRequest
}

type ListEntitlementManagementRoleEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleEligibilityScheduleRequests ...
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequests(ctx context.Context) (result ListEntitlementManagementRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementRoleEligibilityScheduleRequestsCustomPager{},
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
		Values *[]stable.UnifiedRoleEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequestsComplete(ctx context.Context) (ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, UnifiedRoleEligibilityScheduleRequestOperationPredicate{})
}

// ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleEligibilityScheduleRequestOperationPredicate) (result ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleRequest, 0)

	resp, err := c.ListEntitlementManagementRoleEligibilityScheduleRequests(ctx)
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

	result = ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
