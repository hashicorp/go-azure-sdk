package dailyprintusagebyprinter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDailyPrintUsageByPrintersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrintUsageByPrinter
}

type ListDailyPrintUsageByPrintersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrintUsageByPrinter
}

type ListDailyPrintUsageByPrintersOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDailyPrintUsageByPrintersOperationOptions() ListDailyPrintUsageByPrintersOperationOptions {
	return ListDailyPrintUsageByPrintersOperationOptions{}
}

func (o ListDailyPrintUsageByPrintersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDailyPrintUsageByPrintersOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListDailyPrintUsageByPrintersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDailyPrintUsageByPrintersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDailyPrintUsageByPrintersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDailyPrintUsageByPrinters - List dailyPrintUsageByPrinter. Retrieve a list of daily print usage summaries,
// grouped by printer.
func (c DailyPrintUsageByPrinterClient) ListDailyPrintUsageByPrinters(ctx context.Context, options ListDailyPrintUsageByPrintersOperationOptions) (result ListDailyPrintUsageByPrintersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDailyPrintUsageByPrintersCustomPager{},
		Path:          "/reports/dailyPrintUsageByPrinter",
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
		Values *[]beta.PrintUsageByPrinter `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDailyPrintUsageByPrintersComplete retrieves all the results into a single object
func (c DailyPrintUsageByPrinterClient) ListDailyPrintUsageByPrintersComplete(ctx context.Context, options ListDailyPrintUsageByPrintersOperationOptions) (ListDailyPrintUsageByPrintersCompleteResult, error) {
	return c.ListDailyPrintUsageByPrintersCompleteMatchingPredicate(ctx, options, PrintUsageByPrinterOperationPredicate{})
}

// ListDailyPrintUsageByPrintersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DailyPrintUsageByPrinterClient) ListDailyPrintUsageByPrintersCompleteMatchingPredicate(ctx context.Context, options ListDailyPrintUsageByPrintersOperationOptions, predicate PrintUsageByPrinterOperationPredicate) (result ListDailyPrintUsageByPrintersCompleteResult, err error) {
	items := make([]beta.PrintUsageByPrinter, 0)

	resp, err := c.ListDailyPrintUsageByPrinters(ctx, options)
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

	result = ListDailyPrintUsageByPrintersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
