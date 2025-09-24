package relaynamespaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RelayNamespace
}

type NamespacesListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RelayNamespace
}

type NamespacesListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NamespacesListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NamespacesListByResourceGroup ...
func (c RelayNamespacesClient) NamespacesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result NamespacesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NamespacesListByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Relay/namespaces", id.ID()),
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
		Values *[]RelayNamespace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NamespacesListByResourceGroupComplete retrieves all the results into a single object
func (c RelayNamespacesClient) NamespacesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (NamespacesListByResourceGroupCompleteResult, error) {
	return c.NamespacesListByResourceGroupCompleteMatchingPredicate(ctx, id, RelayNamespaceOperationPredicate{})
}

// NamespacesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RelayNamespacesClient) NamespacesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate RelayNamespaceOperationPredicate) (result NamespacesListByResourceGroupCompleteResult, err error) {
	items := make([]RelayNamespace, 0)

	resp, err := c.NamespacesListByResourceGroup(ctx, id)
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

	result = NamespacesListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
