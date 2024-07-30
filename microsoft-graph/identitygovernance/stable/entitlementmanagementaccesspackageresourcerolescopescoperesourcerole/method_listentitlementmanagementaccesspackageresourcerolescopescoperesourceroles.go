package entitlementmanagementaccesspackageresourcerolescopescoperesourcerole

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

type ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoles ...
func (c EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoles(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (result ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCustomPager{},
		Path:       fmt.Sprintf("%s/scope/resource/roles", id.ID()),
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

// ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId) (ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient) ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeId, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoles(ctx, id)
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

	result = ListEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
