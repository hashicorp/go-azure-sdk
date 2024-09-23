package devicehealthscript

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

type GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *GetDeviceHealthScriptGlobalScriptHighestAvailableVersionResult
}

type GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions() GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions {
	return GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions{}
}

func (o GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceHealthScriptGlobalScriptHighestAvailableVersion - Invoke action getGlobalScriptHighestAvailableVersion.
// Update the Proprietary Device Health Script
func (c DeviceHealthScriptClient) GetDeviceHealthScriptGlobalScriptHighestAvailableVersion(ctx context.Context, id beta.DeviceManagementDeviceHealthScriptId, options GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions) (result GetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/getGlobalScriptHighestAvailableVersion", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model GetDeviceHealthScriptGlobalScriptHighestAvailableVersionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
