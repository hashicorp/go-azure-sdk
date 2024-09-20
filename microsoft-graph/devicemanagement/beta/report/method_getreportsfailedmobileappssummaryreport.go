package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsFailedMobileAppsSummaryReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsFailedMobileAppsSummaryReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetReportsFailedMobileAppsSummaryReportOperationOptions() GetReportsFailedMobileAppsSummaryReportOperationOptions {
	return GetReportsFailedMobileAppsSummaryReportOperationOptions{}
}

func (o GetReportsFailedMobileAppsSummaryReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsFailedMobileAppsSummaryReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsFailedMobileAppsSummaryReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsFailedMobileAppsSummaryReport - Invoke action getFailedMobileAppsSummaryReport
func (c ReportClient) GetReportsFailedMobileAppsSummaryReport(ctx context.Context, input GetReportsFailedMobileAppsSummaryReportRequest, options GetReportsFailedMobileAppsSummaryReportOperationOptions) (result GetReportsFailedMobileAppsSummaryReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getFailedMobileAppsSummaryReport",
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
