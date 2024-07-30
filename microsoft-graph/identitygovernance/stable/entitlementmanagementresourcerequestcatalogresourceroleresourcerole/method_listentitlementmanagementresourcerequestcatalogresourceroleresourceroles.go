package entitlementmanagementresourcerequestcatalogresourceroleresourcerole

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

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRoles ...
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRoles(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleId) (result ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCustomPager{},
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

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleId) (ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteMatchingPredicate(ctx, id, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient) ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceRoleId, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRoles(ctx, id)
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

	result = ListEntitlementManagementResourceRequestCatalogResourceRoleResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
