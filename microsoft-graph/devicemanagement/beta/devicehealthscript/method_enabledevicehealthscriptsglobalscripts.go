package devicehealthscript

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableDeviceHealthScriptsGlobalScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EnableDeviceHealthScriptsGlobalScriptsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultEnableDeviceHealthScriptsGlobalScriptsOperationOptions() EnableDeviceHealthScriptsGlobalScriptsOperationOptions {
	return EnableDeviceHealthScriptsGlobalScriptsOperationOptions{}
}

func (o EnableDeviceHealthScriptsGlobalScriptsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EnableDeviceHealthScriptsGlobalScriptsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EnableDeviceHealthScriptsGlobalScriptsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EnableDeviceHealthScriptsGlobalScripts - Invoke action enableGlobalScripts
func (c DeviceHealthScriptClient) EnableDeviceHealthScriptsGlobalScripts(ctx context.Context, options EnableDeviceHealthScriptsGlobalScriptsOperationOptions) (result EnableDeviceHealthScriptsGlobalScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceHealthScripts/enableGlobalScripts",
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

	return
}
