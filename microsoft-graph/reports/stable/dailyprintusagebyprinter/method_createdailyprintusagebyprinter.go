package dailyprintusagebyprinter

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDailyPrintUsageByPrinterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrintUsageByPrinter
}

type CreateDailyPrintUsageByPrinterOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDailyPrintUsageByPrinterOperationOptions() CreateDailyPrintUsageByPrinterOperationOptions {
	return CreateDailyPrintUsageByPrinterOperationOptions{}
}

func (o CreateDailyPrintUsageByPrinterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDailyPrintUsageByPrinterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDailyPrintUsageByPrinterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDailyPrintUsageByPrinter - Create new navigation property to dailyPrintUsageByPrinter for reports
func (c DailyPrintUsageByPrinterClient) CreateDailyPrintUsageByPrinter(ctx context.Context, input stable.PrintUsageByPrinter, options CreateDailyPrintUsageByPrinterOperationOptions) (result CreateDailyPrintUsageByPrinterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/dailyPrintUsageByPrinter",
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

	var model stable.PrintUsageByPrinter
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
