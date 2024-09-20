package devicecustomattributeshellscriptdevicerunstate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions() UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions {
	return UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions{}
}

func (o UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCustomAttributeShellScriptDeviceRunState - Update the navigation property deviceRunStates in
// deviceManagement
func (c DeviceCustomAttributeShellScriptDeviceRunStateClient) UpdateDeviceCustomAttributeShellScriptDeviceRunState(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptIdDeviceRunStateId, input beta.DeviceManagementScriptDeviceState, options UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationOptions) (result UpdateDeviceCustomAttributeShellScriptDeviceRunStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
