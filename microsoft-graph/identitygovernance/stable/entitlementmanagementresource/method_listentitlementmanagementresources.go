package entitlementmanagementresource

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

type ListEntitlementManagementResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResource
}

type ListEntitlementManagementResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResource
}

type ListEntitlementManagementResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResources ...
func (c EntitlementManagementResourceClient) ListEntitlementManagementResources(ctx context.Context) (result ListEntitlementManagementResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourcesCustomPager{},
		Path:       "/identityGovernance/entitlementManagement/resources",
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

// ListEntitlementManagementResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceClient) ListEntitlementManagementResourcesComplete(ctx context.Context) (ListEntitlementManagementResourcesCompleteResult, error) {
	return c.ListEntitlementManagementResourcesCompleteMatchingPredicate(ctx, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceClient) ListEntitlementManagementResourcesCompleteMatchingPredicate(ctx context.Context, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementResourcesCompleteResult, err error) {
	items := make([]stable.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementResources(ctx)
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

	result = ListEntitlementManagementResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
