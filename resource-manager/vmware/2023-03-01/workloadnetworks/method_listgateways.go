package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGatewaysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkGateway
}

type ListGatewaysCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkGateway
}

// ListGateways ...
func (c WorkloadNetworksClient) ListGateways(ctx context.Context, id PrivateCloudId) (result ListGatewaysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/gateways", id.ID()),
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
		Values *[]WorkloadNetworkGateway `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGatewaysComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListGatewaysComplete(ctx context.Context, id PrivateCloudId) (ListGatewaysCompleteResult, error) {
	return c.ListGatewaysCompleteMatchingPredicate(ctx, id, WorkloadNetworkGatewayOperationPredicate{})
}

// ListGatewaysCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListGatewaysCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkGatewayOperationPredicate) (result ListGatewaysCompleteResult, err error) {
	items := make([]WorkloadNetworkGateway, 0)

	resp, err := c.ListGateways(ctx, id)
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

	result = ListGatewaysCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
