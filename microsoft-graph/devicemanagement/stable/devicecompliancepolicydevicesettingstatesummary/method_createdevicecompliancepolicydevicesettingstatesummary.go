package devicecompliancepolicydevicesettingstatesummary

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SettingStateDeviceSummary
}

type CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions() CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions {
	return CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions{}
}

func (o CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyDeviceSettingStateSummary - Create new navigation property to deviceSettingStateSummaries
// for deviceManagement
func (c DeviceCompliancePolicyDeviceSettingStateSummaryClient) CreateDeviceCompliancePolicyDeviceSettingStateSummary(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, input stable.SettingStateDeviceSummary, options CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationOptions) (result CreateDeviceCompliancePolicyDeviceSettingStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceSettingStateSummaries", id.ID()),
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

	var model stable.SettingStateDeviceSummary
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
