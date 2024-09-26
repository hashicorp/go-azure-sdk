package monthlyprintusagebyprinter

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMonthlyPrintUsageByPrinterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMonthlyPrintUsageByPrinterOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateMonthlyPrintUsageByPrinterOperationOptions() UpdateMonthlyPrintUsageByPrinterOperationOptions {
	return UpdateMonthlyPrintUsageByPrinterOperationOptions{}
}

func (o UpdateMonthlyPrintUsageByPrinterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMonthlyPrintUsageByPrinterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMonthlyPrintUsageByPrinterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMonthlyPrintUsageByPrinter - Update the navigation property monthlyPrintUsageByPrinter in reports
func (c MonthlyPrintUsageByPrinterClient) UpdateMonthlyPrintUsageByPrinter(ctx context.Context, id stable.ReportMonthlyPrintUsageByPrinterId, input stable.PrintUsageByPrinter, options UpdateMonthlyPrintUsageByPrinterOperationOptions) (result UpdateMonthlyPrintUsageByPrinterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	return
}
