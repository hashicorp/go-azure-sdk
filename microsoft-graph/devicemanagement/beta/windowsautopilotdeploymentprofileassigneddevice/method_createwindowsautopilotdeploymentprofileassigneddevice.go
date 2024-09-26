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

type CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsAutopilotDeviceIdentity
}

type CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions() CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions {
	return CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions{}
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsAutopilotDeploymentProfileAssignedDevice - Create new navigation property to assignedDevices for
// deviceManagement
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) CreateWindowsAutopilotDeploymentProfileAssignedDevice(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileId, input beta.WindowsAutopilotDeviceIdentity, options CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions) (result CreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/assignedDevices", id.ID()),
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

	var model beta.WindowsAutopilotDeviceIdentity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
