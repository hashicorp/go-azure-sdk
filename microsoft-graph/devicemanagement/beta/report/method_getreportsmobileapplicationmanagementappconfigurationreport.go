package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsMobileApplicationManagementAppConfigurationReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsMobileApplicationManagementAppConfigurationReportOperationOptions() GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions {
	return GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions{}
}

func (o GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsMobileApplicationManagementAppConfigurationReport - Invoke action
// getMobileApplicationManagementAppConfigurationReport
func (c ReportClient) GetReportsMobileApplicationManagementAppConfigurationReport(ctx context.Context, input GetReportsMobileApplicationManagementAppConfigurationReportRequest, options GetReportsMobileApplicationManagementAppConfigurationReportOperationOptions) (result GetReportsMobileApplicationManagementAppConfigurationReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getMobileApplicationManagementAppConfigurationReport",
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
