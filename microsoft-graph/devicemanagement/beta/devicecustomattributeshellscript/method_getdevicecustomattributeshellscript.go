package devicecustomattributeshellscript

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceCustomAttributeShellScriptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceCustomAttributeShellScript
}

type GetDeviceCustomAttributeShellScriptOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDeviceCustomAttributeShellScriptOperationOptions() GetDeviceCustomAttributeShellScriptOperationOptions {
	return GetDeviceCustomAttributeShellScriptOperationOptions{}
}

func (o GetDeviceCustomAttributeShellScriptOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCustomAttributeShellScriptOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceCustomAttributeShellScriptOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCustomAttributeShellScript - Get deviceCustomAttributeShellScripts from deviceManagement. The list of device
// custom attribute shell scripts associated with the tenant.
func (c DeviceCustomAttributeShellScriptClient) GetDeviceCustomAttributeShellScript(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options GetDeviceCustomAttributeShellScriptOperationOptions) (result GetDeviceCustomAttributeShellScriptOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.DeviceCustomAttributeShellScript
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
