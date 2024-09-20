package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsCompliancePolicyNonComplianceReportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsCompliancePolicyNonComplianceReportOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetReportsCompliancePolicyNonComplianceReportOperationOptions() GetReportsCompliancePolicyNonComplianceReportOperationOptions {
	return GetReportsCompliancePolicyNonComplianceReportOperationOptions{}
}

func (o GetReportsCompliancePolicyNonComplianceReportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsCompliancePolicyNonComplianceReportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsCompliancePolicyNonComplianceReportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsCompliancePolicyNonComplianceReport - Invoke action getCompliancePolicyNonComplianceReport. Not yet
// documented
func (c ReportClient) GetReportsCompliancePolicyNonComplianceReport(ctx context.Context, input GetReportsCompliancePolicyNonComplianceReportRequest, options GetReportsCompliancePolicyNonComplianceReportOperationOptions) (result GetReportsCompliancePolicyNonComplianceReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getCompliancePolicyNonComplianceReport",
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
