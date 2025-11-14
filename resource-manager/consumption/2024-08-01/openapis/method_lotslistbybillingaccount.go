package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotsListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LotSummary
}

type LotsListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []LotSummary
}

type LotsListByBillingAccountOperationOptions struct {
	Filter *string
}

func DefaultLotsListByBillingAccountOperationOptions() LotsListByBillingAccountOperationOptions {
	return LotsListByBillingAccountOperationOptions{}
}

func (o LotsListByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o LotsListByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o LotsListByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type LotsListByBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *LotsListByBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// LotsListByBillingAccount ...
func (c OpenapisClient) LotsListByBillingAccount(ctx context.Context, id BillingAccountId, options LotsListByBillingAccountOperationOptions) (result LotsListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &LotsListByBillingAccountCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/lots", id.ID()),
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
		Values *[]LotSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// LotsListByBillingAccountComplete retrieves all the results into a single object
func (c OpenapisClient) LotsListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options LotsListByBillingAccountOperationOptions) (LotsListByBillingAccountCompleteResult, error) {
	return c.LotsListByBillingAccountCompleteMatchingPredicate(ctx, id, options, LotSummaryOperationPredicate{})
}

// LotsListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) LotsListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options LotsListByBillingAccountOperationOptions, predicate LotSummaryOperationPredicate) (result LotsListByBillingAccountCompleteResult, err error) {
	items := make([]LotSummary, 0)

	resp, err := c.LotsListByBillingAccount(ctx, id, options)
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

	result = LotsListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
