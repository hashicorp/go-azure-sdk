package entitlementmanagementresourcerequestcatalogaccesspackage

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

type ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackage
}

type ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackage
}

type ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogAccessPackages ...
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackages(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (result ListEntitlementManagementResourceRequestCatalogAccessPackagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRequestCatalogAccessPackagesCustomPager{},
		Path:       fmt.Sprintf("%s/catalog/accessPackages", id.ID()),
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

// ListEntitlementManagementResourceRequestCatalogAccessPackagesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackagesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate(ctx, id, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogAccessPackageClient) ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult, err error) {
	items := make([]stable.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogAccessPackages(ctx, id)
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

	result = ListEntitlementManagementResourceRequestCatalogAccessPackagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
