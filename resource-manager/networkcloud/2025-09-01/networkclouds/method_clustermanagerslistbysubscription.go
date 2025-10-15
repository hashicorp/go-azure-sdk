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

type ClusterManagersListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterManager
}

type ClusterManagersListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterManager
}

type ClusterManagersListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultClusterManagersListBySubscriptionOperationOptions() ClusterManagersListBySubscriptionOperationOptions {
	return ClusterManagersListBySubscriptionOperationOptions{}
}

func (o ClusterManagersListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ClusterManagersListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ClusterManagersListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ClusterManagersListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ClusterManagersListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ClusterManagersListBySubscription ...
func (c NetworkcloudsClient) ClusterManagersListBySubscription(ctx context.Context, id commonids.SubscriptionId, options ClusterManagersListBySubscriptionOperationOptions) (result ClusterManagersListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ClusterManagersListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/clusterManagers", id.ID()),
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
		Values *[]ClusterManager `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ClusterManagersListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) ClusterManagersListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options ClusterManagersListBySubscriptionOperationOptions) (ClusterManagersListBySubscriptionCompleteResult, error) {
	return c.ClusterManagersListBySubscriptionCompleteMatchingPredicate(ctx, id, options, ClusterManagerOperationPredicate{})
}

// ClusterManagersListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) ClusterManagersListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ClusterManagersListBySubscriptionOperationOptions, predicate ClusterManagerOperationPredicate) (result ClusterManagersListBySubscriptionCompleteResult, err error) {
	items := make([]ClusterManager, 0)

	resp, err := c.ClusterManagersListBySubscription(ctx, id, options)
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

	result = ClusterManagersListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
