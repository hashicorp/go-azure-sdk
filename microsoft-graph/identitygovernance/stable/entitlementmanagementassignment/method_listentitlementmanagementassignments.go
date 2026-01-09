package entitlementmanagementassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageAssignment
}

type ListEntitlementManagementAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageAssignment
}

type ListEntitlementManagementAssignmentsOperationOptions struct {
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

func DefaultListEntitlementManagementAssignmentsOperationOptions() ListEntitlementManagementAssignmentsOperationOptions {
	return ListEntitlementManagementAssignmentsOperationOptions{}
}

func (o ListEntitlementManagementAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignments - List assignments. In Microsoft Entra entitlement management, retrieve a list
// of accessPackageAssignment objects. For directory-wide administrators, the resulting list includes all the
// assignments, current and well as expired, that the caller has access to read, across all catalogs and access
// packages. If the caller is on behalf of a delegated user who is assigned only to catalog-specific delegated
// administrative roles, the request must supply a filter to indicate a specific access package, such as:
// $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'.
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignments(ctx context.Context, options ListEntitlementManagementAssignmentsOperationOptions) (result ListEntitlementManagementAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAssignmentsCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/assignments",
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
		Values *[]stable.AccessPackageAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentsComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignmentsComplete(ctx context.Context, options ListEntitlementManagementAssignmentsOperationOptions) (ListEntitlementManagementAssignmentsCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentsCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentOperationPredicate{})
}

// ListEntitlementManagementAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentClient) ListEntitlementManagementAssignmentsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAssignmentsOperationOptions, predicate AccessPackageAssignmentOperationPredicate) (result ListEntitlementManagementAssignmentsCompleteResult, err error) {
	items := make([]stable.AccessPackageAssignment, 0)

	resp, err := c.ListEntitlementManagementAssignments(ctx, options)
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

	result = ListEntitlementManagementAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
