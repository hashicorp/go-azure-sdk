package deviceshellscriptdevicerunstatemanageddevice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceShellScriptDeviceRunStateManagedDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.ManagedDevice
}

type GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions() GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions {
	return GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions{}
}

func (o GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceShellScriptDeviceRunStateManagedDevice - Get managedDevice from deviceManagement. The managed devices that
// executes the device management script.
func (c DeviceShellScriptDeviceRunStateManagedDeviceClient) GetDeviceShellScriptDeviceRunStateManagedDevice(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdDeviceRunStateId, options GetDeviceShellScriptDeviceRunStateManagedDeviceOperationOptions) (result GetDeviceShellScriptDeviceRunStateManagedDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/managedDevice", id.ID()),
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
