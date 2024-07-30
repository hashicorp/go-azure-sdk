package entitlementmanagementaccesspackageresourcerolescoperoleresourcescope

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

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopes ...
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (result ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCustomPager{},
		Path:       fmt.Sprintf("%s/role/resource/scopes", id.ID()),
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

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopes(ctx, id)
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

	result = ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
