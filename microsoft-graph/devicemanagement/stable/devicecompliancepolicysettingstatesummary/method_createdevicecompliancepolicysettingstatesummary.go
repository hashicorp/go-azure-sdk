package devicecompliancepolicysettingstatesummary

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceCompliancePolicySettingStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceCompliancePolicySettingStateSummary
}

type CreateDeviceCompliancePolicySettingStateSummaryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceCompliancePolicySettingStateSummaryOperationOptions() CreateDeviceCompliancePolicySettingStateSummaryOperationOptions {
	return CreateDeviceCompliancePolicySettingStateSummaryOperationOptions{}
}

func (o CreateDeviceCompliancePolicySettingStateSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicySettingStateSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicySettingStateSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicySettingStateSummary - Create deviceCompliancePolicySettingStateSummary. Create a new
// deviceCompliancePolicySettingStateSummary object.
func (c DeviceCompliancePolicySettingStateSummaryClient) CreateDeviceCompliancePolicySettingStateSummary(ctx context.Context, input stable.DeviceCompliancePolicySettingStateSummary, options CreateDeviceCompliancePolicySettingStateSummaryOperationOptions) (result CreateDeviceCompliancePolicySettingStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceCompliancePolicySettingStateSummaries",
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

	var model stable.DeviceCompliancePolicySettingStateSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
