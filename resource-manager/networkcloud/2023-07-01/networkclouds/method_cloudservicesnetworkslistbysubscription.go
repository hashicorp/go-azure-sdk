package networkclouds

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

type CloudServicesNetworksListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CloudServicesNetwork
}

type CloudServicesNetworksListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CloudServicesNetwork
}

// CloudServicesNetworksListBySubscription ...
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result CloudServicesNetworksListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/cloudServicesNetworks", id.ID()),
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
		Values *[]CloudServicesNetwork `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CloudServicesNetworksListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (CloudServicesNetworksListBySubscriptionCompleteResult, error) {
	return c.CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate(ctx, id, CloudServicesNetworkOperationPredicate{})
}

// CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate CloudServicesNetworkOperationPredicate) (result CloudServicesNetworksListBySubscriptionCompleteResult, err error) {
	items := make([]CloudServicesNetwork, 0)

	resp, err := c.CloudServicesNetworksListBySubscription(ctx, id)
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

	result = CloudServicesNetworksListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
