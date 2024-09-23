package entitlementmanagementroleassignmentscheduleinstance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleInstance
}

type ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleInstance
}

type ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions struct {
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

func DefaultListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions() ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions {
	return ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions{}
}

func (o ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleAssignmentScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleAssignmentScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleAssignmentScheduleInstances - Get roleAssignmentScheduleInstances from roleManagement
func (c EntitlementManagementRoleAssignmentScheduleInstanceClient) ListEntitlementManagementRoleAssignmentScheduleInstances(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions) (result ListEntitlementManagementRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleAssignmentScheduleInstancesCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleAssignmentScheduleInstances",
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
		Values *[]beta.UnifiedRoleAssignmentScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c EntitlementManagementRoleAssignmentScheduleInstanceClient) ListEntitlementManagementRoleAssignmentScheduleInstancesComplete(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions) (ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleInstanceOperationPredicate{})
}

// ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleAssignmentScheduleInstanceClient) ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleInstancesOperationOptions, predicate UnifiedRoleAssignmentScheduleInstanceOperationPredicate) (result ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleInstance, 0)

	resp, err := c.ListEntitlementManagementRoleAssignmentScheduleInstances(ctx, options)
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

	result = ListEntitlementManagementRoleAssignmentScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
