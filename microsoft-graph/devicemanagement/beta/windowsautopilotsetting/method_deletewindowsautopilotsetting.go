package windowsautopilotsetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteWindowsAutopilotSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteWindowsAutopilotSettingOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteWindowsAutopilotSettingOperationOptions() DeleteWindowsAutopilotSettingOperationOptions {
	return DeleteWindowsAutopilotSettingOperationOptions{}
}

func (o DeleteWindowsAutopilotSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteWindowsAutopilotSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteWindowsAutopilotSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteWindowsAutopilotSetting - Delete navigation property windowsAutopilotSettings for deviceManagement
func (c WindowsAutopilotSettingClient) DeleteWindowsAutopilotSetting(ctx context.Context, options DeleteWindowsAutopilotSettingOperationOptions) (result DeleteWindowsAutopilotSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsAutopilotSettings",
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

	return
}
