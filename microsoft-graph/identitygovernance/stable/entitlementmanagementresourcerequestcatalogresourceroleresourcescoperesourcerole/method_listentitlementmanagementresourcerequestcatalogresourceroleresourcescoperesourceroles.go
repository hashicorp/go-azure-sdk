package entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole

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

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions() ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions {
	return ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoles - Get roles from
// identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoles(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, options ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions) (result ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/resource/roles", id.ID()),
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
		Values *[]stable.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, options ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions) (ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleIdResourceScopeId, options ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
