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

type TrunkedNetworksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TrunkedNetwork
}

type TrunkedNetworksListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TrunkedNetwork
}

type TrunkedNetworksListByResourceGroupOperationOptions struct {
	Top *int64
}

func DefaultTrunkedNetworksListByResourceGroupOperationOptions() TrunkedNetworksListByResourceGroupOperationOptions {
	return TrunkedNetworksListByResourceGroupOperationOptions{}
}

func (o TrunkedNetworksListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TrunkedNetworksListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TrunkedNetworksListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type TrunkedNetworksListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TrunkedNetworksListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TrunkedNetworksListByResourceGroup ...
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options TrunkedNetworksListByResourceGroupOperationOptions) (result TrunkedNetworksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &TrunkedNetworksListByResourceGroupCustomPager{},
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

// TrunkedNetworksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options TrunkedNetworksListByResourceGroupOperationOptions) (TrunkedNetworksListByResourceGroupCompleteResult, error) {
	return c.TrunkedNetworksListByResourceGroupCompleteMatchingPredicate(ctx, id, options, TrunkedNetworkOperationPredicate{})
}

// TrunkedNetworksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options TrunkedNetworksListByResourceGroupOperationOptions, predicate TrunkedNetworkOperationPredicate) (result TrunkedNetworksListByResourceGroupCompleteResult, err error) {
	items := make([]TrunkedNetwork, 0)

	resp, err := c.TrunkedNetworksListByResourceGroup(ctx, id, options)
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

	result = TrunkedNetworksListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
