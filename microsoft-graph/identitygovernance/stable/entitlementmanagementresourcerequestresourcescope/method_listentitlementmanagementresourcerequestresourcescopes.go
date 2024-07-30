package entitlementmanagementresourcerequestresourcescope

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

type ListEntitlementManagementResourceRequestResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRequestResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceScope
}

type ListEntitlementManagementResourceRequestResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestResourceScopes ...
func (c EntitlementManagementResourceRequestResourceScopeClient) ListEntitlementManagementResourceRequestResourceScopes(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (result ListEntitlementManagementResourceRequestResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRequestResourceScopesCustomPager{},
		Path:       fmt.Sprintf("%s/resource/scopes", id.ID()),
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
		Values *[]stable.AccessPackageResourceScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRequestResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestResourceScopeClient) ListEntitlementManagementResourceRequestResourceScopesComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (ListEntitlementManagementResourceRequestResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestResourceScopesCompleteMatchingPredicate(ctx, id, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementResourceRequestResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestResourceScopeClient) ListEntitlementManagementResourceRequestResourceScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementResourceRequestResourceScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementResourceRequestResourceScopes(ctx, id)
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

	result = ListEntitlementManagementResourceRequestResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
