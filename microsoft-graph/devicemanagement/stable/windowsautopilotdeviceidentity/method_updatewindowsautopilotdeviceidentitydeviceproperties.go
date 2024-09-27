package windowsautopilotdeviceidentity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions() UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions {
	return UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions{}
}

func (o UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsAutopilotDeviceIdentityDeviceProperties - Invoke action updateDeviceProperties. Updates properties on
// Autopilot devices.
func (c WindowsAutopilotDeviceIdentityClient) UpdateWindowsAutopilotDeviceIdentityDeviceProperties(ctx context.Context, id stable.DeviceManagementWindowsAutopilotDeviceIdentityId, input UpdateWindowsAutopilotDeviceIdentityDevicePropertiesRequest, options UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions) (result UpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/updateDeviceProperties", id.ID()),
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
