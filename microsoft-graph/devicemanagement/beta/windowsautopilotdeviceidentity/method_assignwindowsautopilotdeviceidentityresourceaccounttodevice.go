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

type AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions() AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions {
	return AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions{}
}

func (o AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AssignWindowsAutopilotDeviceIdentityResourceAccountToDevice - Invoke action assignResourceAccountToDevice. Assigns
// resource account to Autopilot devices.
func (c WindowsAutopilotDeviceIdentityClient) AssignWindowsAutopilotDeviceIdentityResourceAccountToDevice(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeviceIdentityId, input AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceRequest, options AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions) (result AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/assignResourceAccountToDevice", id.ID()),
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
