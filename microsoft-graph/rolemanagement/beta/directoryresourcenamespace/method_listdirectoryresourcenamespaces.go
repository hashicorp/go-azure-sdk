package directoryresourcenamespace

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

type ListDirectoryResourceNamespacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceNamespace
}

type ListDirectoryResourceNamespacesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceNamespace
}

type ListDirectoryResourceNamespacesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryResourceNamespacesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryResourceNamespaces ...
func (c DirectoryResourceNamespaceClient) ListDirectoryResourceNamespaces(ctx context.Context) (result ListDirectoryResourceNamespacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryResourceNamespacesCustomPager{},
		Path:       "/roleManagement/directory/resourceNamespaces",
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
		Values *[]beta.UnifiedRbacResourceNamespace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryResourceNamespacesComplete retrieves all the results into a single object
func (c DirectoryResourceNamespaceClient) ListDirectoryResourceNamespacesComplete(ctx context.Context) (ListDirectoryResourceNamespacesCompleteResult, error) {
	return c.ListDirectoryResourceNamespacesCompleteMatchingPredicate(ctx, UnifiedRbacResourceNamespaceOperationPredicate{})
}

// ListDirectoryResourceNamespacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryResourceNamespaceClient) ListDirectoryResourceNamespacesCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRbacResourceNamespaceOperationPredicate) (result ListDirectoryResourceNamespacesCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceNamespace, 0)

	resp, err := c.ListDirectoryResourceNamespaces(ctx)
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

	result = ListDirectoryResourceNamespacesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
