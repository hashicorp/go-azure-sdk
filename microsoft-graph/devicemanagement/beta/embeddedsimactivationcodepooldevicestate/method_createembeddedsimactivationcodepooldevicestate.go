package embeddedsimactivationcodepooldevicestate

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

type CreateEmbeddedSIMActivationCodePoolDeviceStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EmbeddedSIMDeviceState
}

type CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions() CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions {
	return CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions{}
}

func (o CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEmbeddedSIMActivationCodePoolDeviceState - Create new navigation property to deviceStates for deviceManagement
func (c EmbeddedSIMActivationCodePoolDeviceStateClient) CreateEmbeddedSIMActivationCodePoolDeviceState(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, input beta.EmbeddedSIMDeviceState, options CreateEmbeddedSIMActivationCodePoolDeviceStateOperationOptions) (result CreateEmbeddedSIMActivationCodePoolDeviceStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStates", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.EmbeddedSIMDeviceState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
