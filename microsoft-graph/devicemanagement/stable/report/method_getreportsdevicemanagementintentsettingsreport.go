package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsDeviceManagementIntentSettingsReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsDeviceManagementIntentSettingsReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsDeviceManagementIntentSettingsReportOperationOptions() GetReportsDeviceManagementIntentSettingsReportOperationOptions {
	return GetReportsDeviceManagementIntentSettingsReportOperationOptions{}
}

func (o GetReportsDeviceManagementIntentSettingsReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsDeviceManagementIntentSettingsReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsDeviceManagementIntentSettingsReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsDeviceManagementIntentSettingsReport - Invoke action getDeviceManagementIntentSettingsReport. Not yet
// documented
func (c ReportClient) GetReportsDeviceManagementIntentSettingsReport(ctx context.Context, input GetReportsDeviceManagementIntentSettingsReportRequest, options GetReportsDeviceManagementIntentSettingsReportOperationOptions) (result GetReportsDeviceManagementIntentSettingsReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getDeviceManagementIntentSettingsReport",
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
