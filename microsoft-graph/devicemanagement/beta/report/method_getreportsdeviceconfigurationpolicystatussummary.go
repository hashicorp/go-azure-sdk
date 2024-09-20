package report

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsDeviceConfigurationPolicyStatusSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions() GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions {
	return GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions{}
}

func (o GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetReportsDeviceConfigurationPolicyStatusSummary - Invoke action getDeviceConfigurationPolicyStatusSummary
func (c ReportClient) GetReportsDeviceConfigurationPolicyStatusSummary(ctx context.Context, input GetReportsDeviceConfigurationPolicyStatusSummaryRequest, options GetReportsDeviceConfigurationPolicyStatusSummaryOperationOptions) (result GetReportsDeviceConfigurationPolicyStatusSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/reports/getDeviceConfigurationPolicyStatusSummary",
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
