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

type RacksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Rack
}

type RacksListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Rack
}

type RacksListByResourceGroupOperationOptions struct {
	Top *int64
}

func DefaultRacksListByResourceGroupOperationOptions() RacksListByResourceGroupOperationOptions {
	return RacksListByResourceGroupOperationOptions{}
}

func (o RacksListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RacksListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RacksListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type RacksListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RacksListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RacksListByResourceGroup ...
func (c NetworkcloudsClient) RacksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options RacksListByResourceGroupOperationOptions) (result RacksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &RacksListByResourceGroupCustomPager{},
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

// RacksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) RacksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options RacksListByResourceGroupOperationOptions) (RacksListByResourceGroupCompleteResult, error) {
	return c.RacksListByResourceGroupCompleteMatchingPredicate(ctx, id, options, RackOperationPredicate{})
}

// RacksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) RacksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options RacksListByResourceGroupOperationOptions, predicate RackOperationPredicate) (result RacksListByResourceGroupCompleteResult, err error) {
	items := make([]Rack, 0)

	resp, err := c.RacksListByResourceGroup(ctx, id, options)
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

	result = RacksListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
