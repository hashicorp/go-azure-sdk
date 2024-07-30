package directoryresourcenamespaceresourceaction

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

type ListDirectoryResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRbacResourceAction
}

type ListDirectoryResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRbacResourceAction
}

type ListDirectoryResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryResourceNamespaceResourceActions ...
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActions(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId) (result ListDirectoryResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryResourceNamespaceResourceActionsCustomPager{},
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

// ListDirectoryResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActionsComplete(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId) (ListDirectoryResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, UnifiedRbacResourceActionOperationPredicate{})
}

// ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementDirectoryResourceNamespaceId, predicate UnifiedRbacResourceActionOperationPredicate) (result ListDirectoryResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]stable.UnifiedRbacResourceAction, 0)

	resp, err := c.ListDirectoryResourceNamespaceResourceActions(ctx, id)
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

	result = ListDirectoryResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
