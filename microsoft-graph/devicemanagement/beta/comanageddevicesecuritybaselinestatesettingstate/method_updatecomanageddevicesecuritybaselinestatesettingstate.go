package comanageddevicesecuritybaselinestatesettingstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions() UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions {
	return UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions{}
}

func (o UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagedDeviceSecurityBaselineStateSettingState - Update the navigation property settingStates in
// deviceManagement
func (c ComanagedDeviceSecurityBaselineStateSettingStateClient) UpdateComanagedDeviceSecurityBaselineStateSettingState(ctx context.Context, id beta.DeviceManagementComanagedDeviceIdSecurityBaselineStateIdSettingStateId, input beta.SecurityBaselineSettingState, options UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationOptions) (result UpdateComanagedDeviceSecurityBaselineStateSettingStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
