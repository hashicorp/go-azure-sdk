package remoteassistancesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRemoteAssistanceSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateRemoteAssistanceSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateRemoteAssistanceSettingOperationOptions() UpdateRemoteAssistanceSettingOperationOptions {
	return UpdateRemoteAssistanceSettingOperationOptions{}
}

func (o UpdateRemoteAssistanceSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateRemoteAssistanceSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateRemoteAssistanceSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateRemoteAssistanceSetting - Update the navigation property remoteAssistanceSettings in deviceManagement
func (c RemoteAssistanceSettingClient) UpdateRemoteAssistanceSetting(ctx context.Context, input beta.RemoteAssistanceSettings, options UpdateRemoteAssistanceSettingOperationOptions) (result UpdateRemoteAssistanceSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/remoteAssistanceSettings",
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
