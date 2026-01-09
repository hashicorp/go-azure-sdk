package deviceshellscriptuserrunstatedevicerunstatemanageddevice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.ManagedDevice
}

type GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions() GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions {
	return GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions{}
}

func (o GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceShellScriptUserRunStateDeviceRunStateManagedDevice - Get managedDevice from deviceManagement. The managed
// devices that executes the device management script.
func (c DeviceShellScriptUserRunStateDeviceRunStateManagedDeviceClient) GetDeviceShellScriptUserRunStateDeviceRunStateManagedDevice(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdUserRunStateIdDeviceRunStateId, options GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationOptions) (result GetDeviceShellScriptUserRunStateDeviceRunStateManagedDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/managedDevice", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalManagedDeviceImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
