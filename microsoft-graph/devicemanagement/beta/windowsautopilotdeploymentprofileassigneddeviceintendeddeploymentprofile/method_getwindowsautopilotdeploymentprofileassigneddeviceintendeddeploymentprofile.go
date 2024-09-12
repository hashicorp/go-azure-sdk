package windowsautopilotdeploymentprofileassigneddeviceintendeddeploymentprofile

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

type GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.WindowsAutopilotDeploymentProfile
}

type GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions() GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions {
	return GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions{}
}

func (o GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfile - Get intendedDeploymentProfile from
// deviceManagement. Deployment profile intended to be assigned to the Windows autopilot device.
func (c WindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileClient) GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfile(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, options GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationOptions) (result GetWindowsAutopilotDeploymentProfileAssignedDeviceIntendedDeploymentProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/intendedDeploymentProfile", id.ID()),
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
	model, err := beta.UnmarshalWindowsAutopilotDeploymentProfileImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
