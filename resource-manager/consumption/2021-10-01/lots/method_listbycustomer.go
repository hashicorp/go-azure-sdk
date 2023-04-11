package lots

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByCustomerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LotSummary
}

type ListByCustomerCompleteResult struct {
	Items []LotSummary
}

type ListByCustomerOperationOptions struct {
	Filter *string
}

func DefaultListByCustomerOperationOptions() ListByCustomerOperationOptions {
	return ListByCustomerOperationOptions{}
}

func (o ListByCustomerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByCustomerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByCustomerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByCustomer ...
func (c LotsClient) ListByCustomer(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions) (result ListByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/lots", id.ID()),
		OptionsObject: options,
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

// ListByCustomerComplete retrieves all the results into a single object
func (c LotsClient) ListByCustomerComplete(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions) (ListByCustomerCompleteResult, error) {
	return c.ListByCustomerCompleteMatchingPredicate(ctx, id, options, LotSummaryOperationPredicate{})
}

// ListByCustomerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LotsClient) ListByCustomerCompleteMatchingPredicate(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions, predicate LotSummaryOperationPredicate) (result ListByCustomerCompleteResult, err error) {
	items := make([]LotSummary, 0)

	resp, err := c.ListByCustomer(ctx, id, options)
	if err != nil {
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

	result = ListByCustomerCompleteResult{
		Items: items,
	}
	return
}
