package entitlementmanagementcatalogresourcescoperesourcescope

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

type ListEntitlementManagementCatalogResourceScopeResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions struct {
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

func DefaultListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions() ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions {
	return ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions{}
}

func (o ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementCatalogResourceScopeResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourceScopeResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResourceScopeResourceScopes - Get scopes from identityGovernance. Read-only.
// Nullable. Supports $expand.
func (c EntitlementManagementCatalogResourceScopeResourceScopeClient) ListEntitlementManagementCatalogResourceScopeResourceScopes(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId, options ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions) (result ListEntitlementManagementCatalogResourceScopeResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementCatalogResourceScopeResourceScopesCustomPager{},
		Path:          fmt.Sprintf("%s/resource/scopes", id.ID()),
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
		Values *[]stable.AccessPackageResourceScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementCatalogResourceScopeResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceScopeResourceScopeClient) ListEntitlementManagementCatalogResourceScopeResourceScopesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId, options ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions) (ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceScopeResourceScopeClient) ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceScopeId, options ListEntitlementManagementCatalogResourceScopeResourceScopesOperationOptions, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementCatalogResourceScopeResourceScopes(ctx, id, options)
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

	result = ListEntitlementManagementCatalogResourceScopeResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
