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

type BareMetalMachinesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachine
}

type BareMetalMachinesListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BareMetalMachine
}

type BareMetalMachinesListBySubscriptionOperationOptions struct {
	Top *int64
}

func DefaultBareMetalMachinesListBySubscriptionOperationOptions() BareMetalMachinesListBySubscriptionOperationOptions {
	return BareMetalMachinesListBySubscriptionOperationOptions{}
}

func (o BareMetalMachinesListBySubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o BareMetalMachinesListBySubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BareMetalMachinesListBySubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type BareMetalMachinesListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BareMetalMachinesListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BareMetalMachinesListBySubscription ...
func (c NetworkcloudsClient) BareMetalMachinesListBySubscription(ctx context.Context, id commonids.SubscriptionId, options BareMetalMachinesListBySubscriptionOperationOptions) (result BareMetalMachinesListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &BareMetalMachinesListBySubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/bareMetalMachines", id.ID()),
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
		Values *[]BareMetalMachine `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BareMetalMachinesListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BareMetalMachinesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options BareMetalMachinesListBySubscriptionOperationOptions) (BareMetalMachinesListBySubscriptionCompleteResult, error) {
	return c.BareMetalMachinesListBySubscriptionCompleteMatchingPredicate(ctx, id, options, BareMetalMachineOperationPredicate{})
}

// BareMetalMachinesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BareMetalMachinesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options BareMetalMachinesListBySubscriptionOperationOptions, predicate BareMetalMachineOperationPredicate) (result BareMetalMachinesListBySubscriptionCompleteResult, err error) {
	items := make([]BareMetalMachine, 0)

	resp, err := c.BareMetalMachinesListBySubscription(ctx, id, options)
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

	result = BareMetalMachinesListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
