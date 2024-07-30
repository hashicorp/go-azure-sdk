package entitlementmanagementresourcenamespace

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

type ListEntitlementManagementResourceNamespacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRbacResourceNamespace
}

type ListEntitlementManagementResourceNamespacesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRbacResourceNamespace
}

type ListEntitlementManagementResourceNamespacesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceNamespacesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceNamespaces ...
func (c EntitlementManagementResourceNamespaceClient) ListEntitlementManagementResourceNamespaces(ctx context.Context) (result ListEntitlementManagementResourceNamespacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceNamespacesCustomPager{},
		Path:       "/roleManagement/entitlementManagement/resourceNamespaces",
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
		Values *[]stable.UnifiedRbacResourceNamespace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceNamespacesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceNamespaceClient) ListEntitlementManagementResourceNamespacesComplete(ctx context.Context) (ListEntitlementManagementResourceNamespacesCompleteResult, error) {
	return c.ListEntitlementManagementResourceNamespacesCompleteMatchingPredicate(ctx, UnifiedRbacResourceNamespaceOperationPredicate{})
}

// ListEntitlementManagementResourceNamespacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceNamespaceClient) ListEntitlementManagementResourceNamespacesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRbacResourceNamespaceOperationPredicate) (result ListEntitlementManagementResourceNamespacesCompleteResult, err error) {
	items := make([]stable.UnifiedRbacResourceNamespace, 0)

	resp, err := c.ListEntitlementManagementResourceNamespaces(ctx)
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

	result = ListEntitlementManagementResourceNamespacesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
