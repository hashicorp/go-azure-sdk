package entitlementmanagementresourcerolescope

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

type ListEntitlementManagementResourceRoleScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementResourceRoleScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementResourceRoleScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRoleScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRoleScopes ...
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopes(ctx context.Context) (result ListEntitlementManagementResourceRoleScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRoleScopesCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/resourceRoleScopes",
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

// ListEntitlementManagementResourceRoleScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopesComplete(ctx context.Context) (ListEntitlementManagementResourceRoleScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate(ctx, AccessPackageResourceRoleScopeOperationPredicate{})
}

// ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate(ctx context.Context, predicate AccessPackageResourceRoleScopeOperationPredicate) (result ListEntitlementManagementResourceRoleScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRoleScope, 0)

	resp, err := c.ListEntitlementManagementResourceRoleScopes(ctx)
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

	result = ListEntitlementManagementResourceRoleScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
