package entitlementmanagementroleassignmentschedulerequest

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

type ListEntitlementManagementRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions struct {
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

func DefaultListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions() ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions {
	return ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions{}
}

func (o ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleAssignmentScheduleRequests - Get roleAssignmentScheduleRequests from roleManagement
func (c EntitlementManagementRoleAssignmentScheduleRequestClient) ListEntitlementManagementRoleAssignmentScheduleRequests(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions) (result ListEntitlementManagementRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleAssignmentScheduleRequestsCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleAssignmentScheduleRequests",
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
		Values *[]beta.UnifiedRoleAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleAssignmentScheduleRequestClient) ListEntitlementManagementRoleAssignmentScheduleRequestsComplete(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions) (ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleRequestOperationPredicate{})
}

// ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleAssignmentScheduleRequestClient) ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleAssignmentScheduleRequestsOperationOptions, predicate UnifiedRoleAssignmentScheduleRequestOperationPredicate) (result ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleRequest, 0)

	resp, err := c.ListEntitlementManagementRoleAssignmentScheduleRequests(ctx, options)
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

	result = ListEntitlementManagementRoleAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
