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

type UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions() UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions {
	return UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions{}
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperties - Invoke action updateDeviceProperties. Updates
// properties on Autopilot devices.
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperties(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, input UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesRequest, options UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions) (result UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/updateDeviceProperties", id.ID()),
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
