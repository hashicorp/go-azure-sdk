package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions() GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions {
	return GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions{}
}

func (o GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport - Invoke action
// getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (c ReportClient) GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport(ctx context.Context, input GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequest, options GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationOptions) (result GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport",
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
