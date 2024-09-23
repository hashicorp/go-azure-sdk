package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsUnhealthyDefenderAgentsReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsUnhealthyDefenderAgentsReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsUnhealthyDefenderAgentsReportOperationOptions() GetReportsUnhealthyDefenderAgentsReportOperationOptions {
	return GetReportsUnhealthyDefenderAgentsReportOperationOptions{}
}

func (o GetReportsUnhealthyDefenderAgentsReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsUnhealthyDefenderAgentsReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsUnhealthyDefenderAgentsReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsUnhealthyDefenderAgentsReport - Invoke action getUnhealthyDefenderAgentsReport
func (c ReportClient) GetReportsUnhealthyDefenderAgentsReport(ctx context.Context, input GetReportsUnhealthyDefenderAgentsReportRequest, options GetReportsUnhealthyDefenderAgentsReportOperationOptions) (result GetReportsUnhealthyDefenderAgentsReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getUnhealthyDefenderAgentsReport",
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
