package entitlementmanagementresourcescope

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

type ListEntitlementManagementResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceScopesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceScopesOperationOptions() ListEntitlementManagementResourceScopesOperationOptions {
	return ListEntitlementManagementResourceScopesOperationOptions{}
}

func (o ListEntitlementManagementResourceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceScopes - Get scopes from identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementResourceScopeClient) ListEntitlementManagementResourceScopes(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceId, options ListEntitlementManagementResourceScopesOperationOptions) (result ListEntitlementManagementResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceScopesCustomPager{},
		Path:          fmt.Sprintf("%s/scopes", id.ID()),
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

// ListEntitlementManagementResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceScopeClient) ListEntitlementManagementResourceScopesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceId, options ListEntitlementManagementResourceScopesOperationOptions) (ListEntitlementManagementResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceScopeClient) ListEntitlementManagementResourceScopesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceId, options ListEntitlementManagementResourceScopesOperationOptions, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementResourceScopes(ctx, id, options)
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

	result = ListEntitlementManagementResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
