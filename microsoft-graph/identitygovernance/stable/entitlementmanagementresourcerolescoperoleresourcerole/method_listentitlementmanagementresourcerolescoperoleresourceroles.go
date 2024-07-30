package entitlementmanagementresourcerolescoperoleresourcerole

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

type ListEntitlementManagementResourceRoleScopeRoleResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRoleScopeRoleResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRoleScopeRoleResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRoleScopeRoleResourceRoles ...
func (c EntitlementManagementResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementResourceRoleScopeRoleResourceRoles(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId) (result ListEntitlementManagementResourceRoleScopeRoleResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRoleScopeRoleResourceRolesCustomPager{},
		Path:       fmt.Sprintf("%s/role/resource/roles", id.ID()),
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

// ListEntitlementManagementResourceRoleScopeRoleResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementResourceRoleScopeRoleResourceRolesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId) (ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRoleScopeId, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceRoleScopeRoleResourceRoles(ctx, id)
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

	result = ListEntitlementManagementResourceRoleScopeRoleResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
