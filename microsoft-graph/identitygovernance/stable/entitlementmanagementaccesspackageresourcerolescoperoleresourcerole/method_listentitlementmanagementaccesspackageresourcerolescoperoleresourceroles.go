package entitlementmanagementaccesspackageresourcerolescoperoleresourcerole

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

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoles ...
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoles(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (result ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCustomPager{},
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

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoles(ctx, id)
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

	result = ListEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
