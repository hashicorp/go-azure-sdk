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

type ListPublicIPsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkPublicIP
}

type ListPublicIPsCompleteResult struct {
	Items []WorkloadNetworkPublicIP
}

// ListPublicIPs ...
func (c WorkloadNetworksClient) ListPublicIPs(ctx context.Context, id PrivateCloudId) (result ListPublicIPsOperationResponse, err error) {
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

// ListPublicIPsComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListPublicIPsComplete(ctx context.Context, id PrivateCloudId) (ListPublicIPsCompleteResult, error) {
	return c.ListPublicIPsCompleteMatchingPredicate(ctx, id, WorkloadNetworkPublicIPOperationPredicate{})
}

// ListPublicIPsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListPublicIPsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkPublicIPOperationPredicate) (result ListPublicIPsCompleteResult, err error) {
	items := make([]WorkloadNetworkPublicIP, 0)

	resp, err := c.ListPublicIPs(ctx, id)
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

	result = ListPublicIPsCompleteResult{
		Items: items,
	}
	return
}
