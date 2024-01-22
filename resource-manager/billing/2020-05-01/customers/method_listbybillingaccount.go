package customers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Customer
}

type ListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Customer
}

type ListByBillingAccountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultListByBillingAccountOperationOptions() ListByBillingAccountOperationOptions {
	return ListByBillingAccountOperationOptions{}
}

func (o ListByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Search != nil {
		out.Append("$search", fmt.Sprintf("%v", *o.Search))
	}
	return &out
}

// ListByBillingAccount ...
func (c CustomersClient) ListByBillingAccount(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (result ListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/customers", id.ID()),
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
		Values *[]Customer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingAccountComplete retrieves all the results into a single object
func (c CustomersClient) ListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (ListByBillingAccountCompleteResult, error) {
	return c.ListByBillingAccountCompleteMatchingPredicate(ctx, id, options, CustomerOperationPredicate{})
}

// ListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomersClient) ListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions, predicate CustomerOperationPredicate) (result ListByBillingAccountCompleteResult, err error) {
	items := make([]Customer, 0)

	resp, err := c.ListByBillingAccount(ctx, id, options)
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

	result = ListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
