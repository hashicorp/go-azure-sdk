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

type TrunkedNetworksListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TrunkedNetwork
}

type TrunkedNetworksListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TrunkedNetwork
}

type TrunkedNetworksListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultTrunkedNetworksListBySubscriptionOperationOptions() TrunkedNetworksListBySubscriptionOperationOptions {
	return TrunkedNetworksListBySubscriptionOperationOptions{}
}

func (o TrunkedNetworksListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TrunkedNetworksListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TrunkedNetworksListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type TrunkedNetworksListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TrunkedNetworksListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TrunkedNetworksListBySubscription ...
func (c NetworkcloudsClient) TrunkedNetworksListBySubscription(ctx context.Context, id commonids.SubscriptionId, options TrunkedNetworksListBySubscriptionOperationOptions) (result TrunkedNetworksListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &TrunkedNetworksListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/trunkedNetworks", id.ID()),
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
		Values *[]TrunkedNetwork `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TrunkedNetworksListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) TrunkedNetworksListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options TrunkedNetworksListBySubscriptionOperationOptions) (TrunkedNetworksListBySubscriptionCompleteResult, error) {
	return c.TrunkedNetworksListBySubscriptionCompleteMatchingPredicate(ctx, id, options, TrunkedNetworkOperationPredicate{})
}

// TrunkedNetworksListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) TrunkedNetworksListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options TrunkedNetworksListBySubscriptionOperationOptions, predicate TrunkedNetworkOperationPredicate) (result TrunkedNetworksListBySubscriptionCompleteResult, err error) {
	items := make([]TrunkedNetwork, 0)

	resp, err := c.TrunkedNetworksListBySubscription(ctx, id, options)
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

	result = TrunkedNetworksListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
