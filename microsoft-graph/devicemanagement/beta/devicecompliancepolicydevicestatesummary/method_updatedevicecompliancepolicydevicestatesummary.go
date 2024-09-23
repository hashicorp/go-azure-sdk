package devicecompliancepolicydevicestatesummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCompliancePolicyDeviceStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions() UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions {
	return UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyDeviceStateSummary - Update the navigation property
// deviceCompliancePolicyDeviceStateSummary in deviceManagement
func (c DeviceCompliancePolicyDeviceStateSummaryClient) UpdateDeviceCompliancePolicyDeviceStateSummary(ctx context.Context, input beta.DeviceCompliancePolicyDeviceStateSummary, options UpdateDeviceCompliancePolicyDeviceStateSummaryOperationOptions) (result UpdateDeviceCompliancePolicyDeviceStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceCompliancePolicyDeviceStateSummary",
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

	return
}
