package entitlementmanagementaccesspackageassignmentrequest

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

type ListEntitlementManagementAccessPackageAssignmentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageAssignmentRequest
}

type ListEntitlementManagementAccessPackageAssignmentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageAssignmentRequest
}

type ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions() ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentRequests - List accessPackageAssignmentRequests. In Microsoft Entra
// entitlement management, retrieve a list of accessPackageAssignmentRequest objects. The resulting list includes all
// the assignment requests, current and well as expired, that the caller has access to read, across all catalogs and
// access packages.
func (c EntitlementManagementAccessPackageAssignmentRequestClient) ListEntitlementManagementAccessPackageAssignmentRequests(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentRequestsCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackageAssignmentRequests",
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
		Values *[]beta.AccessPackageAssignmentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentRequestsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentRequestClient) ListEntitlementManagementAccessPackageAssignmentRequestsComplete(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions) (ListEntitlementManagementAccessPackageAssignmentRequestsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentRequestsCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentRequestOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentRequestClient) ListEntitlementManagementAccessPackageAssignmentRequestsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentRequestsOperationOptions, predicate AccessPackageAssignmentRequestOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentRequestsCompleteResult, err error) {
	items := make([]beta.AccessPackageAssignmentRequest, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentRequests(ctx, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
