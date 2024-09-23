package entitlementmanagementcatalogresourcerole

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

type ListEntitlementManagementCatalogResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementCatalogResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementCatalogResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementCatalogResourceRolesOperationOptions() ListEntitlementManagementCatalogResourceRolesOperationOptions {
	return ListEntitlementManagementCatalogResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementCatalogResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementCatalogResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementCatalogResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementCatalogResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResourceRoles - Get roles from identityGovernance. Read-only. Nullable. Supports
// $expand.
func (c EntitlementManagementCatalogResourceRoleClient) ListEntitlementManagementCatalogResourceRoles(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceId, options ListEntitlementManagementCatalogResourceRolesOperationOptions) (result ListEntitlementManagementCatalogResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementCatalogResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/roles", id.ID()),
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

// ListEntitlementManagementCatalogResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceRoleClient) ListEntitlementManagementCatalogResourceRolesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceId, options ListEntitlementManagementCatalogResourceRolesOperationOptions) (ListEntitlementManagementCatalogResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementCatalogResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceRoleClient) ListEntitlementManagementCatalogResourceRolesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceId, options ListEntitlementManagementCatalogResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementCatalogResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementCatalogResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementCatalogResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
