package entitlementmanagementaccesspackageassignmentresourcerole

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

type ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageAssignmentResourceRole
}

type ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageAssignmentResourceRole
}

type ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions() ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentResourceRoles - List accessPackageAssignmentResourceRoles. Retrieve a
// list of accessPackageAssignmentResourceRole objects. The resulting list includes all the resource roles of all
// assignments that the caller has access to read, across all catalogs and access packages.
func (c EntitlementManagementAccessPackageAssignmentResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRoles(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentResourceRolesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles",
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
		Values *[]beta.AccessPackageAssignmentResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRolesComplete(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentResourceRoleOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentResourceRolesOperationOptions, predicate AccessPackageAssignmentResourceRoleOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteResult, err error) {
	items := make([]beta.AccessPackageAssignmentResourceRole, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentResourceRoles(ctx, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
