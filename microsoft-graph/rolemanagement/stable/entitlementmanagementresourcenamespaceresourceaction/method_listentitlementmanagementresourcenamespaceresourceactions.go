package entitlementmanagementresourcenamespaceresourceaction

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

type ListEntitlementManagementResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRbacResourceAction
}

type ListEntitlementManagementResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRbacResourceAction
}

type ListEntitlementManagementResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceNamespaceResourceActions ...
func (c EntitlementManagementResourceNamespaceResourceActionClient) ListEntitlementManagementResourceNamespaceResourceActions(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId) (result ListEntitlementManagementResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceNamespaceResourceActionsCustomPager{},
		Path:       fmt.Sprintf("%s/resourceActions", id.ID()),
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
		Values *[]stable.UnifiedRbacResourceAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c EntitlementManagementResourceNamespaceResourceActionClient) ListEntitlementManagementResourceNamespaceResourceActionsComplete(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId) (ListEntitlementManagementResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListEntitlementManagementResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, UnifiedRbacResourceActionOperationPredicate{})
}

// ListEntitlementManagementResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceNamespaceResourceActionClient) ListEntitlementManagementResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEntitlementManagementResourceNamespaceId, predicate UnifiedRbacResourceActionOperationPredicate) (result ListEntitlementManagementResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]stable.UnifiedRbacResourceAction, 0)

	resp, err := c.ListEntitlementManagementResourceNamespaceResourceActions(ctx, id)
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

	result = ListEntitlementManagementResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
