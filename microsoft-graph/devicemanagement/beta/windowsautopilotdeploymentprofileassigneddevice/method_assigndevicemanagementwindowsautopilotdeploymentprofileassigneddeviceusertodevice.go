package windowsautopilotdeploymentprofileassigneddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice ...
func (c WindowsAutopilotDeploymentProfileAssignedDeviceClient) AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice(ctx context.Context, id DeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceId, input AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceRequest) (result AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/assignUserToDevice", id.ID()),
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
