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

type ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions() ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions {
	return ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions{}
}

func (o ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListEntitlementManagementRoleEligibilityScheduleRequests - Get roleEligibilityScheduleRequests from roleManagement.
// Requests for role eligibilities for principals through PIM.
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequests(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions) (result ListEntitlementManagementRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleEligibilityScheduleRequestsCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleEligibilityScheduleRequests",
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
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequestsComplete(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions) (ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleRequestOperationPredicate{})
}

// ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleRequestsOperationOptions, predicate UnifiedRoleEligibilityScheduleRequestOperationPredicate) (result ListEntitlementManagementRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleRequest, 0)

	resp, err := c.ListEntitlementManagementRoleEligibilityScheduleRequests(ctx, options)
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
