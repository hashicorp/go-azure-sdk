package outboundnetworkdependenciesendpoints

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

type ClustersListOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OutboundNetworkDependenciesEndpoint
}

type ClustersListOutboundNetworkDependenciesEndpointsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OutboundNetworkDependenciesEndpoint
}

// ClustersListOutboundNetworkDependenciesEndpoints ...
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpoints(ctx context.Context, id commonids.KustoClusterId) (result ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/outboundNetworkDependenciesEndpoints", id.ID()),
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
		Values *[]OutboundNetworkDependenciesEndpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ClustersListOutboundNetworkDependenciesEndpointsComplete retrieves all the results into a single object
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpointsComplete(ctx context.Context, id commonids.KustoClusterId) (ClustersListOutboundNetworkDependenciesEndpointsCompleteResult, error) {
	return c.ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx, id, OutboundNetworkDependenciesEndpointOperationPredicate{})
}

// ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx context.Context, id commonids.KustoClusterId, predicate OutboundNetworkDependenciesEndpointOperationPredicate) (result ClustersListOutboundNetworkDependenciesEndpointsCompleteResult, err error) {
	items := make([]OutboundNetworkDependenciesEndpoint, 0)

	resp, err := c.ClustersListOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
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

	result = ClustersListOutboundNetworkDependenciesEndpointsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
