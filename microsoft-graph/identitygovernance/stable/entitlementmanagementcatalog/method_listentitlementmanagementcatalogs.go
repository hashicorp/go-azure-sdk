package entitlementmanagementcatalog

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

type ListEntitlementManagementCatalogsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageCatalog
}

type ListEntitlementManagementCatalogsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageCatalog
}

type ListEntitlementManagementCatalogsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementCatalogsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementCatalogs ...
func (c EntitlementManagementCatalogClient) ListEntitlementManagementCatalogs(ctx context.Context) (result ListEntitlementManagementCatalogsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementCatalogsCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/catalogs",
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
		Values *[]stable.AccessPackageCatalog `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementCatalogsComplete retrieves all the results into a single object
func (c EntitlementManagementCatalogClient) ListEntitlementManagementCatalogsComplete(ctx context.Context) (ListEntitlementManagementCatalogsCompleteResult, error) {
	return c.ListEntitlementManagementCatalogsCompleteMatchingPredicate(ctx, AccessPackageCatalogOperationPredicate{})
}

// ListEntitlementManagementCatalogsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementCatalogClient) ListEntitlementManagementCatalogsCompleteMatchingPredicate(ctx context.Context, predicate AccessPackageCatalogOperationPredicate) (result ListEntitlementManagementCatalogsCompleteResult, err error) {
	items := make([]stable.AccessPackageCatalog, 0)

	resp, err := c.ListEntitlementManagementCatalogs(ctx)
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

	result = ListEntitlementManagementCatalogsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
