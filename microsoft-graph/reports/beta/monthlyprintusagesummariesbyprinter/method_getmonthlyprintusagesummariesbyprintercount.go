package monthlyprintusagesummariesbyprinter

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMonthlyPrintUsageSummariesByPrinterCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetMonthlyPrintUsageSummariesByPrinterCountOperationOptions() GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions {
	return GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions{}
}

func (o GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMonthlyPrintUsageSummariesByPrinterCount - Get the number of the resource
func (c MonthlyPrintUsageSummariesByPrinterClient) GetMonthlyPrintUsageSummariesByPrinterCount(ctx context.Context, options GetMonthlyPrintUsageSummariesByPrinterCountOperationOptions) (result GetMonthlyPrintUsageSummariesByPrinterCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/monthlyPrintUsageSummariesByPrinter/$count",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
