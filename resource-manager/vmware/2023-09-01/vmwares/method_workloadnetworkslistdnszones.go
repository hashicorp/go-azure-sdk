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

type WorkloadNetworksListDnsZonesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkDnsZone
}

type WorkloadNetworksListDnsZonesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkDnsZone
}

// WorkloadNetworksListDnsZones ...
func (c VMwaresClient) WorkloadNetworksListDnsZones(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListDnsZonesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/dnsZones", id.ID()),
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
		Values *[]WorkloadNetworkDnsZone `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkloadNetworksListDnsZonesComplete retrieves all the results into a single object
func (c VMwaresClient) WorkloadNetworksListDnsZonesComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListDnsZonesCompleteResult, error) {
	return c.WorkloadNetworksListDnsZonesCompleteMatchingPredicate(ctx, id, WorkloadNetworkDnsZoneOperationPredicate{})
}

// WorkloadNetworksListDnsZonesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMwaresClient) WorkloadNetworksListDnsZonesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDnsZoneOperationPredicate) (result WorkloadNetworksListDnsZonesCompleteResult, err error) {
	items := make([]WorkloadNetworkDnsZone, 0)

	resp, err := c.WorkloadNetworksListDnsZones(ctx, id)
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

	result = WorkloadNetworksListDnsZonesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
