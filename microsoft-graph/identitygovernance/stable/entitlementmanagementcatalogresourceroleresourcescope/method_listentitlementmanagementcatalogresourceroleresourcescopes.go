package entitlementmanagementcatalogresourceroleresourcescope

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

type ListEntitlementManagementCatalogResourceRoleResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceRoleResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourceRoleResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResourceRoleResourceScopes ...
func (c EntitlementManagementCatalogResourceRoleResourceScopeClient) ListEntitlementManagementCatalogResourceRoleResourceScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) (result ListEntitlementManagementCatalogResourceRoleResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementCatalogResourceRoleResourceScopesCustomPager{},
		Path:       fmt.Sprintf("%s/resource/scopes", id.ID()),
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

// ListEntitlementManagementCatalogResourceRoleResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceRoleResourceScopeClient) ListEntitlementManagementCatalogResourceRoleResourceScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId) (ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceRoleResourceScopeClient) ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceRoleId, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementCatalogResourceRoleResourceScopes(ctx, id)
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

	result = ListEntitlementManagementCatalogResourceRoleResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
