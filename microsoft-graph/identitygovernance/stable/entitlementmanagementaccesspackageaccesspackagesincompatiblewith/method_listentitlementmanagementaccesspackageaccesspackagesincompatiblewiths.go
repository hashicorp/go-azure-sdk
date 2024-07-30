package entitlementmanagementaccesspackageaccesspackagesincompatiblewith

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

type ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackage
}

type ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackage
}

type ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWiths ...
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWiths(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (result ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCustomPager{},
		Path:       fmt.Sprintf("%s/accessPackagesIncompatibleWith", id.ID()),
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

// ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId) (ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteMatchingPredicate(ctx, id, AccessPackageOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementAccessPackageId, predicate AccessPackageOperationPredicate) (result ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteResult, err error) {
	items := make([]stable.AccessPackage, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWiths(ctx, id)
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

	result = ListEntitlementManagementAccessPackageAccessPackagesIncompatibleWithsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
