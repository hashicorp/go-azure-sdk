package entitlementmanagementresourcerequestcatalogresourcescoperesourcerole

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

type ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions() ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions {
	return ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRoles - Get roles from identityGovernance.
// Read-only. Nullable. Supports $expand.
func (c EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRoles(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, options ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions) (result ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/resource/roles", id.ID()),
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
		Values *[]stable.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, options ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions) (ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, options ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementResourceRequestCatalogResourceScopeResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
