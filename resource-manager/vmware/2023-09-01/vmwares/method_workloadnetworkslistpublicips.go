package vmwares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksListPublicIPsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkPublicIP
}

type WorkloadNetworksListPublicIPsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkPublicIP
}

// WorkloadNetworksListPublicIPs ...
func (c VMwaresClient) WorkloadNetworksListPublicIPs(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListPublicIPsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/publicIPs", id.ID()),
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
		Values *[]WorkloadNetworkPublicIP `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkloadNetworksListPublicIPsComplete retrieves all the results into a single object
func (c VMwaresClient) WorkloadNetworksListPublicIPsComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListPublicIPsCompleteResult, error) {
	return c.WorkloadNetworksListPublicIPsCompleteMatchingPredicate(ctx, id, WorkloadNetworkPublicIPOperationPredicate{})
}

// WorkloadNetworksListPublicIPsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMwaresClient) WorkloadNetworksListPublicIPsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkPublicIPOperationPredicate) (result WorkloadNetworksListPublicIPsCompleteResult, err error) {
	items := make([]WorkloadNetworkPublicIP, 0)

	resp, err := c.WorkloadNetworksListPublicIPs(ctx, id)
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

	result = WorkloadNetworksListPublicIPsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
