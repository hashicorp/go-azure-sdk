package androidforworksetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAndroidForWorkSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAndroidForWorkSettingOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteAndroidForWorkSettingOperationOptions() DeleteAndroidForWorkSettingOperationOptions {
	return DeleteAndroidForWorkSettingOperationOptions{}
}

func (o DeleteAndroidForWorkSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAndroidForWorkSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteAndroidForWorkSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAndroidForWorkSetting - Delete navigation property androidForWorkSettings for deviceManagement
func (c AndroidForWorkSettingClient) DeleteAndroidForWorkSetting(ctx context.Context, options DeleteAndroidForWorkSettingOperationOptions) (result DeleteAndroidForWorkSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/androidForWorkSettings",
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
