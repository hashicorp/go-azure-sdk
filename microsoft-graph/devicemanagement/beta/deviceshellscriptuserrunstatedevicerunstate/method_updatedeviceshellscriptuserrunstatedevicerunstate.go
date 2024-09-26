package deviceshellscriptuserrunstatedevicerunstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions() UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions {
	return UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions{}
}

func (o UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceShellScriptUserRunStateDeviceRunState - Update the navigation property deviceRunStates in
// deviceManagement
func (c DeviceShellScriptUserRunStateDeviceRunStateClient) UpdateDeviceShellScriptUserRunStateDeviceRunState(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId, input beta.DeviceManagementScriptDeviceState, options UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationOptions) (result UpdateDeviceShellScriptUserRunStateDeviceRunStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
