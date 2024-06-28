package invoice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Invoice
}

type ListByBillingSubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Invoice
}

type ListByBillingSubscriptionOperationOptions struct {
	Count           *bool
	Filter          *string
	OrderBy         *string
	PeriodEndDate   *string
	PeriodStartDate *string
	Search          *string
	Skip            *int64
	Top             *int64
}

func DefaultListByBillingSubscriptionOperationOptions() ListByBillingSubscriptionOperationOptions {
	return ListByBillingSubscriptionOperationOptions{}
}

func (o ListByBillingSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingSubscriptionOperationOptions) ToQuery() *client.QueryParams {
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
	if o.PeriodEndDate != nil {
		out.Append("periodEndDate", fmt.Sprintf("%v", *o.PeriodEndDate))
	}
	if o.PeriodStartDate != nil {
		out.Append("periodStartDate", fmt.Sprintf("%v", *o.PeriodStartDate))
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

type ListByBillingSubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByBillingSubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByBillingSubscription ...
func (c InvoiceClient) ListByBillingSubscription(ctx context.Context, id BillingSubscriptionId, options ListByBillingSubscriptionOperationOptions) (result ListByBillingSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ListByBillingSubscriptionCustomPager{},
		Path:          fmt.Sprintf("%s/invoices", id.ID()),
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
		Values *[]Invoice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingSubscriptionComplete retrieves all the results into a single object
func (c InvoiceClient) ListByBillingSubscriptionComplete(ctx context.Context, id BillingSubscriptionId, options ListByBillingSubscriptionOperationOptions) (ListByBillingSubscriptionCompleteResult, error) {
	return c.ListByBillingSubscriptionCompleteMatchingPredicate(ctx, id, options, InvoiceOperationPredicate{})
}

// ListByBillingSubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InvoiceClient) ListByBillingSubscriptionCompleteMatchingPredicate(ctx context.Context, id BillingSubscriptionId, options ListByBillingSubscriptionOperationOptions, predicate InvoiceOperationPredicate) (result ListByBillingSubscriptionCompleteResult, err error) {
	items := make([]Invoice, 0)

	resp, err := c.ListByBillingSubscription(ctx, id, options)
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

	result = ListByBillingSubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
