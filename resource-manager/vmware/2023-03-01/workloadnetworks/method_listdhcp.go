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

type ListDhcpOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkDhcp
}

type ListDhcpCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkDhcp
}

// ListDhcp ...
func (c WorkloadNetworksClient) ListDhcp(ctx context.Context, id PrivateCloudId) (result ListDhcpOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/dhcpConfigurations", id.ID()),
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
		Values *[]WorkloadNetworkDhcp `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDhcpComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListDhcpComplete(ctx context.Context, id PrivateCloudId) (ListDhcpCompleteResult, error) {
	return c.ListDhcpCompleteMatchingPredicate(ctx, id, WorkloadNetworkDhcpOperationPredicate{})
}

// ListDhcpCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListDhcpCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDhcpOperationPredicate) (result ListDhcpCompleteResult, err error) {
	items := make([]WorkloadNetworkDhcp, 0)

	resp, err := c.ListDhcp(ctx, id)
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

	result = ListDhcpCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
