package openapis

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

type DnsResolversListByVirtualNetworkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubResource
}

type DnsResolversListByVirtualNetworkCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubResource
}

type DnsResolversListByVirtualNetworkOperationOptions struct {
	Top *int64
}

func DefaultDnsResolversListByVirtualNetworkOperationOptions() DnsResolversListByVirtualNetworkOperationOptions {
	return DnsResolversListByVirtualNetworkOperationOptions{}
}

func (o DnsResolversListByVirtualNetworkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DnsResolversListByVirtualNetworkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DnsResolversListByVirtualNetworkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type DnsResolversListByVirtualNetworkCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DnsResolversListByVirtualNetworkCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DnsResolversListByVirtualNetwork ...
func (c OpenapisClient) DnsResolversListByVirtualNetwork(ctx context.Context, id commonids.VirtualNetworkId, options DnsResolversListByVirtualNetworkOperationOptions) (result DnsResolversListByVirtualNetworkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &DnsResolversListByVirtualNetworkCustomPager{},
		Path:          fmt.Sprintf("%s/listDnsResolvers", id.ID()),
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
		Values *[]SubResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DnsResolversListByVirtualNetworkComplete retrieves all the results into a single object
func (c OpenapisClient) DnsResolversListByVirtualNetworkComplete(ctx context.Context, id commonids.VirtualNetworkId, options DnsResolversListByVirtualNetworkOperationOptions) (DnsResolversListByVirtualNetworkCompleteResult, error) {
	return c.DnsResolversListByVirtualNetworkCompleteMatchingPredicate(ctx, id, options, SubResourceOperationPredicate{})
}

// DnsResolversListByVirtualNetworkCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) DnsResolversListByVirtualNetworkCompleteMatchingPredicate(ctx context.Context, id commonids.VirtualNetworkId, options DnsResolversListByVirtualNetworkOperationOptions, predicate SubResourceOperationPredicate) (result DnsResolversListByVirtualNetworkCompleteResult, err error) {
	items := make([]SubResource, 0)

	resp, err := c.DnsResolversListByVirtualNetwork(ctx, id, options)
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

	result = DnsResolversListByVirtualNetworkCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
