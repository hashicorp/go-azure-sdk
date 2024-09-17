package transaction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInvoiceSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Transaction
}

type ListByInvoiceSectionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Transaction
}

type ListByInvoiceSectionOperationOptions struct {
	Count           *bool
	Filter          *string
	OrderBy         *string
	PeriodEndDate   *string
	PeriodStartDate *string
	Search          *string
	Skip            *int64
	Top             *int64
	Type            *TransactionType
}

func DefaultListByInvoiceSectionOperationOptions() ListByInvoiceSectionOperationOptions {
	return ListByInvoiceSectionOperationOptions{}
}

func (o ListByInvoiceSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByInvoiceSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByInvoiceSectionOperationOptions) ToQuery() *client.QueryParams {
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
	if o.Type != nil {
		out.Append("type", fmt.Sprintf("%v", *o.Type))
	}
	return &out
}

type ListByInvoiceSectionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByInvoiceSectionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByInvoiceSection ...
func (c TransactionClient) ListByInvoiceSection(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions) (result ListByInvoiceSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByInvoiceSectionCustomPager{},
		Path:          fmt.Sprintf("%s/transactions", id.ID()),
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
		Values *[]Transaction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInvoiceSectionComplete retrieves all the results into a single object
func (c TransactionClient) ListByInvoiceSectionComplete(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions) (ListByInvoiceSectionCompleteResult, error) {
	return c.ListByInvoiceSectionCompleteMatchingPredicate(ctx, id, options, TransactionOperationPredicate{})
}

// ListByInvoiceSectionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransactionClient) ListByInvoiceSectionCompleteMatchingPredicate(ctx context.Context, id InvoiceSectionId, options ListByInvoiceSectionOperationOptions, predicate TransactionOperationPredicate) (result ListByInvoiceSectionCompleteResult, err error) {
	items := make([]Transaction, 0)

	resp, err := c.ListByInvoiceSection(ctx, id, options)
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

	result = ListByInvoiceSectionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
