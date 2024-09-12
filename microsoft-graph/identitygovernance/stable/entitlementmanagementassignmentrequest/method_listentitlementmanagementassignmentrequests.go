package entitlementmanagementassignmentrequest

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

type ListEntitlementManagementAssignmentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageAssignmentRequest
}

type ListEntitlementManagementAssignmentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageAssignmentRequest
}

type ListEntitlementManagementAssignmentRequestsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAssignmentRequestsOperationOptions() ListEntitlementManagementAssignmentRequestsOperationOptions {
	return ListEntitlementManagementAssignmentRequestsOperationOptions{}
}

func (o ListEntitlementManagementAssignmentRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAssignmentRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAssignmentRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAssignmentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentRequests - List assignmentRequests. In Microsoft Entra entitlement management,
// retrieve a list of accessPackageAssignmentRequest objects. The resulting list includes all the assignment requests,
// current and well as expired, that the caller has access to read, across all catalogs and access packages.
func (c EntitlementManagementAssignmentRequestClient) ListEntitlementManagementAssignmentRequests(ctx context.Context, options ListEntitlementManagementAssignmentRequestsOperationOptions) (result ListEntitlementManagementAssignmentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAssignmentRequestsCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/assignmentRequests",
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
		Values *[]stable.AccessPackageAssignmentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentRequestsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentRequestClient) ListEntitlementManagementAssignmentRequestsComplete(ctx context.Context, options ListEntitlementManagementAssignmentRequestsOperationOptions) (ListEntitlementManagementAssignmentRequestsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentRequestsCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentRequestOperationPredicate{})
}

// ListEntitlementManagementAssignmentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentRequestClient) ListEntitlementManagementAssignmentRequestsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAssignmentRequestsOperationOptions, predicate AccessPackageAssignmentRequestOperationPredicate) (result ListEntitlementManagementAssignmentRequestsCompleteResult, err error) {
	items := make([]stable.AccessPackageAssignmentRequest, 0)

	resp, err := c.ListEntitlementManagementAssignmentRequests(ctx, options)
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

	result = ListEntitlementManagementAssignmentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
