package monthlyprintusagesummariesbyprinter

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMonthlyPrintUsageSummariesByPrinterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PrintUsageByPrinter
}

type CreateMonthlyPrintUsageSummariesByPrinterOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateMonthlyPrintUsageSummariesByPrinterOperationOptions() CreateMonthlyPrintUsageSummariesByPrinterOperationOptions {
	return CreateMonthlyPrintUsageSummariesByPrinterOperationOptions{}
}

func (o CreateMonthlyPrintUsageSummariesByPrinterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMonthlyPrintUsageSummariesByPrinterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMonthlyPrintUsageSummariesByPrinterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMonthlyPrintUsageSummariesByPrinter - Create new navigation property to monthlyPrintUsageSummariesByPrinter for
// reports
func (c MonthlyPrintUsageSummariesByPrinterClient) CreateMonthlyPrintUsageSummariesByPrinter(ctx context.Context, input beta.PrintUsageByPrinter, options CreateMonthlyPrintUsageSummariesByPrinterOperationOptions) (result CreateMonthlyPrintUsageSummariesByPrinterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/monthlyPrintUsageSummariesByPrinter",
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

	var model beta.PrintUsageByPrinter
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
