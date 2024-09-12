package monthlyprintusagebyprinter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMonthlyPrintUsageByPrintersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrintUsageByPrinter
}

type ListMonthlyPrintUsageByPrintersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrintUsageByPrinter
}

type ListMonthlyPrintUsageByPrintersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMonthlyPrintUsageByPrintersOperationOptions() ListMonthlyPrintUsageByPrintersOperationOptions {
	return ListMonthlyPrintUsageByPrintersOperationOptions{}
}

func (o ListMonthlyPrintUsageByPrintersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMonthlyPrintUsageByPrintersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListMonthlyPrintUsageByPrintersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMonthlyPrintUsageByPrintersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonthlyPrintUsageByPrintersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonthlyPrintUsageByPrinters - List monthlyPrintUsageByPrinter. Retrieve a list of monthly print usage summaries,
// grouped by printer.
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrinters(ctx context.Context, options ListMonthlyPrintUsageByPrintersOperationOptions) (result ListMonthlyPrintUsageByPrintersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMonthlyPrintUsageByPrintersCustomPager{},
		Path:          "/reports/monthlyPrintUsageByPrinter",
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
		Values *[]stable.PrintUsageByPrinter `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMonthlyPrintUsageByPrintersComplete retrieves all the results into a single object
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrintersComplete(ctx context.Context, options ListMonthlyPrintUsageByPrintersOperationOptions) (ListMonthlyPrintUsageByPrintersCompleteResult, error) {
	return c.ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate(ctx, options, PrintUsageByPrinterOperationPredicate{})
}

// ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate(ctx context.Context, options ListMonthlyPrintUsageByPrintersOperationOptions, predicate PrintUsageByPrinterOperationPredicate) (result ListMonthlyPrintUsageByPrintersCompleteResult, err error) {
	items := make([]stable.PrintUsageByPrinter, 0)

	resp, err := c.ListMonthlyPrintUsageByPrinters(ctx, options)
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

	result = ListMonthlyPrintUsageByPrintersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
