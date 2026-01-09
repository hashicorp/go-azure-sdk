package entitlementmanagementroleeligibilityschedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementRoleEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleEligibilitySchedule
}

type ListEntitlementManagementRoleEligibilitySchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleEligibilitySchedule
}

type ListEntitlementManagementRoleEligibilitySchedulesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementRoleEligibilitySchedulesOperationOptions() ListEntitlementManagementRoleEligibilitySchedulesOperationOptions {
	return ListEntitlementManagementRoleEligibilitySchedulesOperationOptions{}
}

func (o ListEntitlementManagementRoleEligibilitySchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleEligibilitySchedulesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleEligibilitySchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListEntitlementManagementRoleEligibilitySchedules - Get roleEligibilitySchedules from roleManagement
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedules(ctx context.Context, options ListEntitlementManagementRoleEligibilitySchedulesOperationOptions) (result ListEntitlementManagementRoleEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleEligibilitySchedulesCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleEligibilitySchedules",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.UnifiedRoleEligibilitySchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleEligibilitySchedulesComplete retrieves all the results into a single object
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedulesComplete(ctx context.Context, options ListEntitlementManagementRoleEligibilitySchedulesOperationOptions) (ListEntitlementManagementRoleEligibilitySchedulesCompleteResult, error) {
	return c.ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleOperationPredicate{})
}

// ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleEligibilityScheduleClient) ListEntitlementManagementRoleEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleEligibilitySchedulesOperationOptions, predicate UnifiedRoleEligibilityScheduleOperationPredicate) (result ListEntitlementManagementRoleEligibilitySchedulesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleEligibilitySchedule, 0)

	resp, err := c.ListEntitlementManagementRoleEligibilitySchedules(ctx, options)
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
