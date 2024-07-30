package entitlementmanagementcatalogaccesspackage

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

type ListEntitlementManagementCatalogAccessPackagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackage
}

type ListEntitlementManagementCatalogAccessPackagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackage
}

type ListEntitlementManagementCatalogAccessPackagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogAccessPackagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogAccessPackages ...
func (c EntitlementManagementCatalogAccessPackageClient) ListEntitlementManagementCatalogAccessPackages(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId) (result ListEntitlementManagementCatalogAccessPackagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementCatalogAccessPackagesCustomPager{},
		Path:       fmt.Sprintf("%s/accessPackages", id.ID()),
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
		Values *[]stable.AccessPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementCatalogAccessPackagesComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogAccessPackageClient) ListEntitlementManagementCatalogAccessPackagesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId) (ListEntitlementManagementCatalogAccessPackagesCompleteResult, error) {
	return c.ListEntitlementManagementCatalogAccessPackagesCompleteMatchingPredicate(ctx, id, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementCatalogAccessPackagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogAccessPackageClient) ListEntitlementManagementCatalogAccessPackagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementCatalogId, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementCatalogAccessPackagesCompleteResult, err error) {
	items := make([]stable.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementCatalogAccessPackages(ctx, id)
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

	result = ListEntitlementManagementCatalogAccessPackagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
