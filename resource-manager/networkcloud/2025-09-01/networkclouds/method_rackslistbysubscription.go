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

type RacksListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Rack
}

type RacksListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Rack
}

type RacksListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultRacksListBySubscriptionOperationOptions() RacksListBySubscriptionOperationOptions {
	return RacksListBySubscriptionOperationOptions{}
}

func (o RacksListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RacksListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RacksListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type RacksListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RacksListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RacksListBySubscription ...
func (c NetworkcloudsClient) RacksListBySubscription(ctx context.Context, id commonids.SubscriptionId, options RacksListBySubscriptionOperationOptions) (result RacksListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RacksListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/racks", id.ID()),
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
		Values *[]Rack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RacksListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) RacksListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options RacksListBySubscriptionOperationOptions) (RacksListBySubscriptionCompleteResult, error) {
	return c.RacksListBySubscriptionCompleteMatchingPredicate(ctx, id, options, RackOperationPredicate{})
}

// RacksListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) RacksListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options RacksListBySubscriptionOperationOptions, predicate RackOperationPredicate) (result RacksListBySubscriptionCompleteResult, err error) {
	items := make([]Rack, 0)

	resp, err := c.RacksListBySubscription(ctx, id, options)
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

	result = RacksListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
