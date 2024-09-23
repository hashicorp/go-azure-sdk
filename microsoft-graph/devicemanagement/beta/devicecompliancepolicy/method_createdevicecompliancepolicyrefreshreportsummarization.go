package devicecompliancepolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceCompliancePolicyRefreshReportSummarizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions() CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions {
	return CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions{}
}

func (o CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyRefreshReportSummarization - Invoke action refreshDeviceComplianceReportSummarization
func (c DeviceCompliancePolicyClient) CreateDeviceCompliancePolicyRefreshReportSummarization(ctx context.Context, options CreateDeviceCompliancePolicyRefreshReportSummarizationOperationOptions) (result CreateDeviceCompliancePolicyRefreshReportSummarizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceCompliancePolicies/refreshDeviceComplianceReportSummarization",
		RetryFunc:     options.RetryFunc,
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

	return
}
