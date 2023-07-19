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

type ListDnsServicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkDnsService
}

type ListDnsServicesCompleteResult struct {
	Items []WorkloadNetworkDnsService
}

// ListDnsServices ...
func (c WorkloadNetworksClient) ListDnsServices(ctx context.Context, id PrivateCloudId) (result ListDnsServicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/dnsServices", id.ID()),
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
		Values *[]WorkloadNetworkDnsService `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDnsServicesComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListDnsServicesComplete(ctx context.Context, id PrivateCloudId) (ListDnsServicesCompleteResult, error) {
	return c.ListDnsServicesCompleteMatchingPredicate(ctx, id, WorkloadNetworkDnsServiceOperationPredicate{})
}

// ListDnsServicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListDnsServicesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDnsServiceOperationPredicate) (result ListDnsServicesCompleteResult, err error) {
	items := make([]WorkloadNetworkDnsService, 0)

	resp, err := c.ListDnsServices(ctx, id)
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

	result = ListDnsServicesCompleteResult{
		Items: items,
	}
	return
}
