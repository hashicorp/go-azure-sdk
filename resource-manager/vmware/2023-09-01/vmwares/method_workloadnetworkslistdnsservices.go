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

type WorkloadNetworksListDnsServicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkDnsService
}

type WorkloadNetworksListDnsServicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkDnsService
}

// WorkloadNetworksListDnsServices ...
func (c VMwaresClient) WorkloadNetworksListDnsServices(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListDnsServicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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

// WorkloadNetworksListDnsServicesComplete retrieves all the results into a single object
func (c VMwaresClient) WorkloadNetworksListDnsServicesComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListDnsServicesCompleteResult, error) {
	return c.WorkloadNetworksListDnsServicesCompleteMatchingPredicate(ctx, id, WorkloadNetworkDnsServiceOperationPredicate{})
}

// WorkloadNetworksListDnsServicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMwaresClient) WorkloadNetworksListDnsServicesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDnsServiceOperationPredicate) (result WorkloadNetworksListDnsServicesCompleteResult, err error) {
	items := make([]WorkloadNetworkDnsService, 0)

	resp, err := c.WorkloadNetworksListDnsServices(ctx, id)
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

	result = WorkloadNetworksListDnsServicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
