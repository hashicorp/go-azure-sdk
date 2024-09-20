package entitlementmanagementtransitiveroleassignment

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

type ListEntitlementManagementTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignment
}

type ListEntitlementManagementTransitiveRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignment
}

type ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListEntitlementManagementTransitiveRoleAssignmentsOperationOptions() ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions {
	return ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions{}
}

func (o ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementTransitiveRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementTransitiveRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementTransitiveRoleAssignments - Get transitiveRoleAssignments from roleManagement
func (c EntitlementManagementTransitiveRoleAssignmentClient) ListEntitlementManagementTransitiveRoleAssignments(ctx context.Context, options ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions) (result ListEntitlementManagementTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementTransitiveRoleAssignmentsCustomPager{},
		Path:          "/roleManagement/entitlementManagement/transitiveRoleAssignments",
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
		Values *[]beta.UnifiedRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c EntitlementManagementTransitiveRoleAssignmentClient) ListEntitlementManagementTransitiveRoleAssignmentsComplete(ctx context.Context, options ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions) (ListEntitlementManagementTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentOperationPredicate{})
}

// ListEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementTransitiveRoleAssignmentClient) ListEntitlementManagementTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementTransitiveRoleAssignmentsOperationOptions, predicate UnifiedRoleAssignmentOperationPredicate) (result ListEntitlementManagementTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignment, 0)

	resp, err := c.ListEntitlementManagementTransitiveRoleAssignments(ctx, options)
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

	result = ListEntitlementManagementTransitiveRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
