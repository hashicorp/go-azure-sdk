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

type L2NetworksListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]L2Network
}

type L2NetworksListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []L2Network
}

type L2NetworksListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultL2NetworksListBySubscriptionOperationOptions() L2NetworksListBySubscriptionOperationOptions {
	return L2NetworksListBySubscriptionOperationOptions{}
}

func (o L2NetworksListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o L2NetworksListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o L2NetworksListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type L2NetworksListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *L2NetworksListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// L2NetworksListBySubscription ...
func (c NetworkcloudsClient) L2NetworksListBySubscription(ctx context.Context, id commonids.SubscriptionId, options L2NetworksListBySubscriptionOperationOptions) (result L2NetworksListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &L2NetworksListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/l2Networks", id.ID()),
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
		Values *[]L2Network `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// L2NetworksListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) L2NetworksListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options L2NetworksListBySubscriptionOperationOptions) (L2NetworksListBySubscriptionCompleteResult, error) {
	return c.L2NetworksListBySubscriptionCompleteMatchingPredicate(ctx, id, options, L2NetworkOperationPredicate{})
}

// L2NetworksListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) L2NetworksListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options L2NetworksListBySubscriptionOperationOptions, predicate L2NetworkOperationPredicate) (result L2NetworksListBySubscriptionCompleteResult, err error) {
	items := make([]L2Network, 0)

	resp, err := c.L2NetworksListBySubscription(ctx, id, options)
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

	result = L2NetworksListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
