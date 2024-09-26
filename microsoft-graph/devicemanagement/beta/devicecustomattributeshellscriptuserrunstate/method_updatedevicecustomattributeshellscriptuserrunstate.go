package devicecustomattributeshellscriptuserrunstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCustomAttributeShellScriptUserRunStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions() UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions {
	return UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions{}
}

func (o UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCustomAttributeShellScriptUserRunState - Update the navigation property userRunStates in deviceManagement
func (c DeviceCustomAttributeShellScriptUserRunStateClient) UpdateDeviceCustomAttributeShellScriptUserRunState(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptIdUserRunStateId, input beta.DeviceManagementScriptUserState, options UpdateDeviceCustomAttributeShellScriptUserRunStateOperationOptions) (result UpdateDeviceCustomAttributeShellScriptUserRunStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
