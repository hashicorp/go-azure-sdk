package entitlementmanagementresourceenvironmentresourcescoperesourcerole

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

type ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceEnvironmentResourceScopeResourceRoles ...
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceRoleClient) ListEntitlementManagementResourceEnvironmentResourceScopeResourceRoles(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) (result ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCustomPager{},
		Path:       fmt.Sprintf("%s/resource/roles", id.ID()),
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

// ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceRoleClient) ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId) (ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceRoleClient) ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceEnvironmentResourceScopeResourceRoles(ctx, id)
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

	result = ListEntitlementManagementResourceEnvironmentResourceScopeResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
