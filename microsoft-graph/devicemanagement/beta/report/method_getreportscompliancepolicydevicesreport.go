package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsCompliancePolicyDevicesReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsCompliancePolicyDevicesReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetReportsCompliancePolicyDevicesReportOperationOptions() GetReportsCompliancePolicyDevicesReportOperationOptions {
	return GetReportsCompliancePolicyDevicesReportOperationOptions{}
}

func (o GetReportsCompliancePolicyDevicesReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsCompliancePolicyDevicesReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsCompliancePolicyDevicesReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsCompliancePolicyDevicesReport - Invoke action getCompliancePolicyDevicesReport
func (c ReportClient) GetReportsCompliancePolicyDevicesReport(ctx context.Context, input GetReportsCompliancePolicyDevicesReportRequest, options GetReportsCompliancePolicyDevicesReportOperationOptions) (result GetReportsCompliancePolicyDevicesReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getCompliancePolicyDevicesReport",
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
