package embeddedsimactivationcodepooldevicestate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEmbeddedSIMActivationCodePoolDeviceStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EmbeddedSIMDeviceState
}

type GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions() GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions {
	return GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions{}
}

func (o GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToOData() *odata.Query {
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

func (o GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEmbeddedSIMActivationCodePoolDeviceState - Get deviceStates from deviceManagement. Navigational property to a list
// of device states for this pool.
func (c EmbeddedSIMActivationCodePoolDeviceStateClient) GetEmbeddedSIMActivationCodePoolDeviceState(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolIdDeviceStateId, options GetEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) (result GetEmbeddedSIMActivationCodePoolDeviceStateOperationResponse, err error) {
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

	var model beta.EmbeddedSIMDeviceState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
