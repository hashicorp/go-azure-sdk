package entitlementmanagementresourcerequestresourceroleresourcescope

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

type ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions() ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions {
	return ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions{}
}

func (o ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceRequestResourceRoleResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestResourceRoleResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestResourceRoleResourceScopes - Get scopes from identityGovernance. Read-only.
// Nullable. Supports $expand.
func (c EntitlementManagementResourceRequestResourceRoleResourceScopeClient) ListEntitlementManagementResourceRequestResourceRoleResourceScopes(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, options ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions) (result ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceRequestResourceRoleResourceScopesCustomPager{},
		Path:          fmt.Sprintf("%s/resource/scopes", id.ID()),
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

// ListEntitlementManagementResourceRequestResourceRoleResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestResourceRoleResourceScopeClient) ListEntitlementManagementResourceRequestResourceRoleResourceScopesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, options ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions) (ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestResourceRoleResourceScopeClient) ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, options ListEntitlementManagementResourceRequestResourceRoleResourceScopesOperationOptions, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementResourceRequestResourceRoleResourceScopes(ctx, id, options)
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

	result = ListEntitlementManagementResourceRequestResourceRoleResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
