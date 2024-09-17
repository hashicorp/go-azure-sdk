package billingrequest

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
	Model        *[]BillingRequest
}

type ListByCustomerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRequest
}

type ListByCustomerOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *string
	Search  *string
	Skip    *int64
	Top     *int64
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
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByCustomerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByCustomerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCustomer ...
func (c BillingRequestClient) ListByCustomer(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions) (result ListByCustomerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByCustomerCustomPager{},
		Path:          fmt.Sprintf("%s/billingRequests", id.ID()),
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
		Values *[]BillingRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByCustomerComplete retrieves all the results into a single object
func (c BillingRequestClient) ListByCustomerComplete(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions) (ListByCustomerCompleteResult, error) {
	return c.ListByCustomerCompleteMatchingPredicate(ctx, id, options, BillingRequestOperationPredicate{})
}

// ListByCustomerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRequestClient) ListByCustomerCompleteMatchingPredicate(ctx context.Context, id BillingProfileCustomerId, options ListByCustomerOperationOptions, predicate BillingRequestOperationPredicate) (result ListByCustomerCompleteResult, err error) {
	items := make([]BillingRequest, 0)

	resp, err := c.ListByCustomer(ctx, id, options)
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

	result = ListByCustomerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
