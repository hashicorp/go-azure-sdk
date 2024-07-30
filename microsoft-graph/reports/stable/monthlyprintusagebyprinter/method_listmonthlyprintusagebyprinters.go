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

type ListMonthlyPrintUsageByPrintersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonthlyPrintUsageByPrintersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonthlyPrintUsageByPrinters ...
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrinters(ctx context.Context) (result ListMonthlyPrintUsageByPrintersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMonthlyPrintUsageByPrintersCustomPager{},
		Path:       "/reports/monthlyPrintUsageByPrinter",
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
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrintersComplete(ctx context.Context) (ListMonthlyPrintUsageByPrintersCompleteResult, error) {
	return c.ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate(ctx, PrintUsageByPrinterOperationPredicate{})
}

// ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonthlyPrintUsageByPrinterClient) ListMonthlyPrintUsageByPrintersCompleteMatchingPredicate(ctx context.Context, predicate PrintUsageByPrinterOperationPredicate) (result ListMonthlyPrintUsageByPrintersCompleteResult, err error) {
	items := make([]stable.PrintUsageByPrinter, 0)

	resp, err := c.ListMonthlyPrintUsageByPrinters(ctx)
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
