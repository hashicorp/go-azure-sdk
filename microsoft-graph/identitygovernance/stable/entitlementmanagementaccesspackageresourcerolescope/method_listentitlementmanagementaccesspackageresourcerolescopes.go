package entitlementmanagementaccesspackageresourcerolescope

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

type ListEntitlementManagementAccessPackageResourceRoleScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageResourceRoleScopes ...
func (c EntitlementManagementAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (result ListEntitlementManagementAccessPackageResourceRoleScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageResourceRoleScopesCustomPager{},
		Path:       fmt.Sprintf("%s/resourceRoleScopes", id.ID()),
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
		Values *[]stable.AccessPackageResourceRoleScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageResourceRoleScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageResourceRoleScopeClient) ListEntitlementManagementAccessPackageResourceRoleScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId, predicate AccessPackageResourceRoleScopeOperationPredicate) (result ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRoleScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageResourceRoleScopes(ctx, id)
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

	result = ListEntitlementManagementAccessPackageResourceRoleScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
