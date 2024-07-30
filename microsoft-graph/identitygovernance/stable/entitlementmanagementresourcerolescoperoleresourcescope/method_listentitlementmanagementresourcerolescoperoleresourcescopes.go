package entitlementmanagementresourcerolescoperoleresourcescope

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

type ListEntitlementManagementResourceRoleScopeRoleResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRoleScopeRoleResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRoleScopeRoleResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRoleScopeRoleResourceScopes ...
func (c EntitlementManagementResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementResourceRoleScopeRoleResourceScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId) (result ListEntitlementManagementResourceRoleScopeRoleResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRoleScopeRoleResourceScopesCustomPager{},
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

// ListEntitlementManagementResourceRoleScopeRoleResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementResourceRoleScopeRoleResourceScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId) (ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRoleScopeRoleResourceScopeClient) ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementResourceRoleScopeRoleResourceScopes(ctx, id)
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

	result = ListEntitlementManagementResourceRoleScopeRoleResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
