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

type DnsForwardingRulesetsListByVirtualNetworkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualNetworkDnsForwardingRuleset
}

type DnsForwardingRulesetsListByVirtualNetworkCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VirtualNetworkDnsForwardingRuleset
}

type DnsForwardingRulesetsListByVirtualNetworkOperationOptions struct {
	Top *int64
}

func DefaultDnsForwardingRulesetsListByVirtualNetworkOperationOptions() DnsForwardingRulesetsListByVirtualNetworkOperationOptions {
	return DnsForwardingRulesetsListByVirtualNetworkOperationOptions{}
}

func (o DnsForwardingRulesetsListByVirtualNetworkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DnsForwardingRulesetsListByVirtualNetworkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DnsForwardingRulesetsListByVirtualNetworkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type DnsForwardingRulesetsListByVirtualNetworkCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DnsForwardingRulesetsListByVirtualNetworkCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DnsForwardingRulesetsListByVirtualNetwork ...
func (c OpenapisClient) DnsForwardingRulesetsListByVirtualNetwork(ctx context.Context, id commonids.VirtualNetworkId, options DnsForwardingRulesetsListByVirtualNetworkOperationOptions) (result DnsForwardingRulesetsListByVirtualNetworkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &DnsForwardingRulesetsListByVirtualNetworkCustomPager{},
		Path:          fmt.Sprintf("%s/listDnsForwardingRulesets", id.ID()),
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
		Values *[]VirtualNetworkDnsForwardingRuleset `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DnsForwardingRulesetsListByVirtualNetworkComplete retrieves all the results into a single object
func (c OpenapisClient) DnsForwardingRulesetsListByVirtualNetworkComplete(ctx context.Context, id commonids.VirtualNetworkId, options DnsForwardingRulesetsListByVirtualNetworkOperationOptions) (DnsForwardingRulesetsListByVirtualNetworkCompleteResult, error) {
	return c.DnsForwardingRulesetsListByVirtualNetworkCompleteMatchingPredicate(ctx, id, options, VirtualNetworkDnsForwardingRulesetOperationPredicate{})
}

// DnsForwardingRulesetsListByVirtualNetworkCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) DnsForwardingRulesetsListByVirtualNetworkCompleteMatchingPredicate(ctx context.Context, id commonids.VirtualNetworkId, options DnsForwardingRulesetsListByVirtualNetworkOperationOptions, predicate VirtualNetworkDnsForwardingRulesetOperationPredicate) (result DnsForwardingRulesetsListByVirtualNetworkCompleteResult, err error) {
	items := make([]VirtualNetworkDnsForwardingRuleset, 0)

	resp, err := c.DnsForwardingRulesetsListByVirtualNetwork(ctx, id, options)
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

	result = DnsForwardingRulesetsListByVirtualNetworkCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
