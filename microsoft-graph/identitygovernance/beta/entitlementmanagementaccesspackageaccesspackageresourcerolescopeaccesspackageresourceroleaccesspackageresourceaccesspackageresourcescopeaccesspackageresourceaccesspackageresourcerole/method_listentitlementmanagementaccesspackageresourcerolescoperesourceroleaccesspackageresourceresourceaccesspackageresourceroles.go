package entitlementmanagementaccesspackageaccesspackageresourcerolescopeaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

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

type ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions() ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions {
	return ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRoles
// - Get accessPackageResourceRoles from identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRoles(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions) (result ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResource/accessPackageResourceRoles", id.ID()),
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
		Values *[]beta.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions) (ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeIdAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeId, options ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageResourceRoleScopeResourceRoleAccessPackageResourceResourceAccessPackageResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
