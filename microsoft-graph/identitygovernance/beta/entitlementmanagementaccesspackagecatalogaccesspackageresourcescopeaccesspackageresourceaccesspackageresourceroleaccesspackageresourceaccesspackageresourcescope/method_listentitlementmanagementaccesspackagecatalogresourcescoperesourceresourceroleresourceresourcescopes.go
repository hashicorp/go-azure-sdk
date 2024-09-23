package entitlementmanagementaccesspackagecatalogaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourcescope

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

type ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions() ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions {
	return ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopes - Get
// accessPackageResourceScopes from identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopes(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, options ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions) (result ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResource/accessPackageResourceScopes", id.ID()),
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
		Values *[]beta.AccessPackageResourceScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, options ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions) (ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, options ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesOperationOptions, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopes(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageCatalogResourceScopeResourceResourceRoleResourceResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
