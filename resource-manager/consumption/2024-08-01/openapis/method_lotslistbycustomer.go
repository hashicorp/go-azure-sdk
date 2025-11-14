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

type LotsListByCustomerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LotSummary
}

type LotsListByCustomerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []LotSummary
}

type LotsListByCustomerOperationOptions struct {
	Filter *string
}

func DefaultLotsListByCustomerOperationOptions() LotsListByCustomerOperationOptions {
	return LotsListByCustomerOperationOptions{}
}

func (o LotsListByCustomerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o LotsListByCustomerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o LotsListByCustomerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type LotsListByCustomerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *LotsListByCustomerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// LotsListByCustomer ...
func (c OpenapisClient) LotsListByCustomer(ctx context.Context, id CustomerId, options LotsListByCustomerOperationOptions) (result LotsListByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &LotsListByCustomerCustomPager{},
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

// LotsListByCustomerComplete retrieves all the results into a single object
func (c OpenapisClient) LotsListByCustomerComplete(ctx context.Context, id CustomerId, options LotsListByCustomerOperationOptions) (LotsListByCustomerCompleteResult, error) {
	return c.LotsListByCustomerCompleteMatchingPredicate(ctx, id, options, LotSummaryOperationPredicate{})
}

// LotsListByCustomerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) LotsListByCustomerCompleteMatchingPredicate(ctx context.Context, id CustomerId, options LotsListByCustomerOperationOptions, predicate LotSummaryOperationPredicate) (result LotsListByCustomerCompleteResult, err error) {
	items := make([]LotSummary, 0)

	resp, err := c.LotsListByCustomer(ctx, id, options)
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

	result = LotsListByCustomerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
