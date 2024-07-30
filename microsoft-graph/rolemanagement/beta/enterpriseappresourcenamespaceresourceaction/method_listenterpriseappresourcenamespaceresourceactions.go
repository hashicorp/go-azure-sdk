package enterpriseappresourcenamespaceresourceaction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEnterpriseAppResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceAction
}

type ListEnterpriseAppResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceAction
}

type ListEnterpriseAppResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppResourceNamespaceResourceActions ...
func (c EnterpriseAppResourceNamespaceResourceActionClient) ListEnterpriseAppResourceNamespaceResourceActions(ctx context.Context, id RoleManagementEnterpriseAppIdResourceNamespaceId) (result ListEnterpriseAppResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppResourceNamespaceResourceActionsCustomPager{},
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
		Values *[]beta.UnifiedRbacResourceAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c EnterpriseAppResourceNamespaceResourceActionClient) ListEnterpriseAppResourceNamespaceResourceActionsComplete(ctx context.Context, id RoleManagementEnterpriseAppIdResourceNamespaceId) (ListEnterpriseAppResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListEnterpriseAppResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, UnifiedRbacResourceActionOperationPredicate{})
}

// ListEnterpriseAppResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppResourceNamespaceResourceActionClient) ListEnterpriseAppResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppIdResourceNamespaceId, predicate UnifiedRbacResourceActionOperationPredicate) (result ListEnterpriseAppResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceAction, 0)

	resp, err := c.ListEnterpriseAppResourceNamespaceResourceActions(ctx, id)
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

	result = ListEnterpriseAppResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
