package workloadnetworkgateways

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksListGatewaysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkGateway
}

type WorkloadNetworksListGatewaysCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkGateway
}

type WorkloadNetworksListGatewaysCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkloadNetworksListGatewaysCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkloadNetworksListGateways ...
func (c WorkloadNetworkGatewaysClient) WorkloadNetworksListGateways(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListGatewaysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkloadNetworksListGatewaysCustomPager{},
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

// WorkloadNetworksListGatewaysComplete retrieves all the results into a single object
func (c WorkloadNetworkGatewaysClient) WorkloadNetworksListGatewaysComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListGatewaysCompleteResult, error) {
	return c.WorkloadNetworksListGatewaysCompleteMatchingPredicate(ctx, id, WorkloadNetworkGatewayOperationPredicate{})
}

// WorkloadNetworksListGatewaysCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworkGatewaysClient) WorkloadNetworksListGatewaysCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkGatewayOperationPredicate) (result WorkloadNetworksListGatewaysCompleteResult, err error) {
	items := make([]WorkloadNetworkGateway, 0)

	resp, err := c.WorkloadNetworksListGateways(ctx, id)
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

	result = WorkloadNetworksListGatewaysCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
