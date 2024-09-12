package entitlementmanagementroleeligibilityscheduleinstance

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

type ListEntitlementManagementRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleInstance
}

type ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleInstance
}

type ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions() ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions {
	return ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions{}
}

func (o ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleEligibilityScheduleInstances - Get roleEligibilityScheduleInstances from roleManagement.
// Instances for role eligibility requests.
func (c EntitlementManagementRoleEligibilityScheduleInstanceClient) ListEntitlementManagementRoleEligibilityScheduleInstances(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions) (result ListEntitlementManagementRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleEligibilityScheduleInstancesCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleEligibilityScheduleInstances",
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
		Values *[]stable.UnifiedRoleEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c EntitlementManagementRoleEligibilityScheduleInstanceClient) ListEntitlementManagementRoleEligibilityScheduleInstancesComplete(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions) (ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleInstanceOperationPredicate{})
}

// ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleEligibilityScheduleInstanceClient) ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleEligibilityScheduleInstancesOperationOptions, predicate UnifiedRoleEligibilityScheduleInstanceOperationPredicate) (result ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleInstance, 0)

	resp, err := c.ListEntitlementManagementRoleEligibilityScheduleInstances(ctx, options)
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

	result = ListEntitlementManagementRoleEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
