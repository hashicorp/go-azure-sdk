package devicecustomattributeshellscriptuserrunstate

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

type GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions() GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions {
	return GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions{}
}

func (o GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCustomAttributeShellScriptUserRunStatesCount - Get the number of the resource
func (c DeviceCustomAttributeShellScriptUserRunStateClient) GetDeviceCustomAttributeShellScriptUserRunStatesCount(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationOptions) (result GetDeviceCustomAttributeShellScriptUserRunStatesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userRunStates/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
