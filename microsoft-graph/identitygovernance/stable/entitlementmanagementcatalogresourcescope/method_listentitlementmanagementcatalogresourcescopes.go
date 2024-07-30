package entitlementmanagementcatalogresourcescope

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

type ListEntitlementManagementCatalogResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementCatalogResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResourceScopes ...
func (c EntitlementManagementCatalogResourceScopeClient) ListEntitlementManagementCatalogResourceScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceId) (result ListEntitlementManagementCatalogResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementCatalogResourceScopesCustomPager{},
		Path:       fmt.Sprintf("%s/scopes", id.ID()),
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

// ListEntitlementManagementCatalogResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceScopeClient) ListEntitlementManagementCatalogResourceScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceId) (ListEntitlementManagementCatalogResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourceScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementCatalogResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceScopeClient) ListEntitlementManagementCatalogResourceScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogIdResourceId, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementCatalogResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementCatalogResourceScopes(ctx, id)
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

	result = ListEntitlementManagementCatalogResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
