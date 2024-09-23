package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions() GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions {
	return GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions{}
}

func (o GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsDevicesStatusByPolicyPlatformComplianceReport - Invoke action
// getDevicesStatusByPolicyPlatformComplianceReport
func (c ReportClient) GetReportsDevicesStatusByPolicyPlatformComplianceReport(ctx context.Context, input GetReportsDevicesStatusByPolicyPlatformComplianceReportRequest, options GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationOptions) (result GetReportsDevicesStatusByPolicyPlatformComplianceReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getDevicesStatusByPolicyPlatformComplianceReport",
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
