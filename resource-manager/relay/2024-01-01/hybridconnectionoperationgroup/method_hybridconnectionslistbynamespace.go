package hybridconnectionoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionsListByNamespaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HybridConnection
}

type HybridConnectionsListByNamespaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HybridConnection
}

type HybridConnectionsListByNamespaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *HybridConnectionsListByNamespaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// HybridConnectionsListByNamespace ...
func (c HybridConnectionOperationGroupClient) HybridConnectionsListByNamespace(ctx context.Context, id NamespaceId) (result HybridConnectionsListByNamespaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &HybridConnectionsListByNamespaceCustomPager{},
		Path:       fmt.Sprintf("%s/hybridConnections", id.ID()),
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
		Values *[]HybridConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// HybridConnectionsListByNamespaceComplete retrieves all the results into a single object
func (c HybridConnectionOperationGroupClient) HybridConnectionsListByNamespaceComplete(ctx context.Context, id NamespaceId) (HybridConnectionsListByNamespaceCompleteResult, error) {
	return c.HybridConnectionsListByNamespaceCompleteMatchingPredicate(ctx, id, HybridConnectionOperationPredicate{})
}

// HybridConnectionsListByNamespaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HybridConnectionOperationGroupClient) HybridConnectionsListByNamespaceCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate HybridConnectionOperationPredicate) (result HybridConnectionsListByNamespaceCompleteResult, err error) {
	items := make([]HybridConnection, 0)

	resp, err := c.HybridConnectionsListByNamespace(ctx, id)
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

	result = HybridConnectionsListByNamespaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
