package androidforworksetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAndroidForWorkSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAndroidForWorkSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAndroidForWorkSettingOperationOptions() UpdateAndroidForWorkSettingOperationOptions {
	return UpdateAndroidForWorkSettingOperationOptions{}
}

func (o UpdateAndroidForWorkSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAndroidForWorkSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAndroidForWorkSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAndroidForWorkSetting - Update the navigation property androidForWorkSettings in deviceManagement
func (c AndroidForWorkSettingClient) UpdateAndroidForWorkSetting(ctx context.Context, input beta.AndroidForWorkSettings, options UpdateAndroidForWorkSettingOperationOptions) (result UpdateAndroidForWorkSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/androidForWorkSettings",
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
