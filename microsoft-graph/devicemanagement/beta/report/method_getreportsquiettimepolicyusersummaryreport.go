package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsQuietTimePolicyUserSummaryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsQuietTimePolicyUserSummaryReportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetReportsQuietTimePolicyUserSummaryReportOperationOptions() GetReportsQuietTimePolicyUserSummaryReportOperationOptions {
	return GetReportsQuietTimePolicyUserSummaryReportOperationOptions{}
}

func (o GetReportsQuietTimePolicyUserSummaryReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsQuietTimePolicyUserSummaryReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsQuietTimePolicyUserSummaryReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsQuietTimePolicyUserSummaryReport - Invoke action getQuietTimePolicyUserSummaryReport
func (c ReportClient) GetReportsQuietTimePolicyUserSummaryReport(ctx context.Context, input GetReportsQuietTimePolicyUserSummaryReportRequest, options GetReportsQuietTimePolicyUserSummaryReportOperationOptions) (result GetReportsQuietTimePolicyUserSummaryReportOperationResponse, err error) {
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
		Path:          "/deviceManagement/reports/getQuietTimePolicyUserSummaryReport",
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
