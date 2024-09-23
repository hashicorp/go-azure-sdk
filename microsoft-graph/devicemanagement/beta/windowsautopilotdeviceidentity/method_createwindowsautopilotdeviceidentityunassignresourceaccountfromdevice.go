package windowsautopilotdeviceidentity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions() CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions {
	return CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions{}
}

func (o CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDevice - Invoke action
// unassignResourceAccountFromDevice. Unassigns the resource account from an Autopilot device.
func (c WindowsAutopilotDeviceIdentityClient) CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDevice(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeviceIdentityId, options CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions) (result CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unassignResourceAccountFromDevice", id.ID()),
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

	return
}
