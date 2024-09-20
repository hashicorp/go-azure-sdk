package windowsautopilotdeploymentprofileassigneddevice

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

type CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions() CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions {
	return CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions{}
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDevice - Invoke action
// unassignResourceAccountFromDevice. Unassigns the resource account from an Autopilot device.
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDevice(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, options CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions) (result CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unassignResourceAccountFromDevice", id.ID()),
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
