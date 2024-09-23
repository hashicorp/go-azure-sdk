package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsWindowsDriverUpdateAlertSummaryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions() GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions {
	return GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions{}
}

func (o GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsWindowsDriverUpdateAlertSummaryReport - Invoke action getWindowsDriverUpdateAlertSummaryReport
func (c ReportClient) GetReportsWindowsDriverUpdateAlertSummaryReport(ctx context.Context, input GetReportsWindowsDriverUpdateAlertSummaryReportRequest, options GetReportsWindowsDriverUpdateAlertSummaryReportOperationOptions) (result GetReportsWindowsDriverUpdateAlertSummaryReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getWindowsDriverUpdateAlertSummaryReport",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
