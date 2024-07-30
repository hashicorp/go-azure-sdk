package entitlementmanagementcatalogresource

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

type ListEntitlementManagementCatalogResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResource
}

type ListEntitlementManagementCatalogResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResource
}

type ListEntitlementManagementCatalogResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogResources ...
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResources(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId) (result ListEntitlementManagementCatalogResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementCatalogResourcesCustomPager{},
		Path:       fmt.Sprintf("%s/resources", id.ID()),
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
		Values *[]stable.AccessPackageResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementCatalogResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResourcesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId) (ListEntitlementManagementCatalogResourcesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate(ctx, id, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogResourceClient) ListEntitlementManagementCatalogResourcesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementCatalogResourcesCompleteResult, err error) {
	items := make([]stable.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementCatalogResources(ctx, id)
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

	result = ListEntitlementManagementCatalogResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
