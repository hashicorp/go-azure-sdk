package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageresourcerolescope

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

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions() ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopes - Get accessPackageResourceRoleScopes
// from identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopes(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackage/accessPackageResourceRoleScopes", id.ID()),
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
		Values *[]beta.AccessPackageResourceRoleScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesOperationOptions, predicate AccessPackageResourceRoleScopeOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceRoleScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopes(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
