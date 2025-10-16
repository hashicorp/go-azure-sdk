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

type CloudServicesNetworksListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultCloudServicesNetworksListBySubscriptionOperationOptions() CloudServicesNetworksListBySubscriptionOperationOptions {
	return CloudServicesNetworksListBySubscriptionOperationOptions{}
}

func (o CloudServicesNetworksListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CloudServicesNetworksListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CloudServicesNetworksListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type CloudServicesNetworksListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *CloudServicesNetworksListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CloudServicesNetworksListBySubscription ...
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscription(ctx context.Context, id commonids.SubscriptionId, options CloudServicesNetworksListBySubscriptionOperationOptions) (result CloudServicesNetworksListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &CloudServicesNetworksListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/cloudServicesNetworks", id.ID()),
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
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options CloudServicesNetworksListBySubscriptionOperationOptions) (CloudServicesNetworksListBySubscriptionCompleteResult, error) {
	return c.CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate(ctx, id, options, CloudServicesNetworkOperationPredicate{})
}

// CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) CloudServicesNetworksListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options CloudServicesNetworksListBySubscriptionOperationOptions, predicate CloudServicesNetworkOperationPredicate) (result CloudServicesNetworksListBySubscriptionCompleteResult, err error) {
	items := make([]CloudServicesNetwork, 0)

	resp, err := c.CloudServicesNetworksListBySubscription(ctx, id, options)
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

	result = CloudServicesNetworksListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
