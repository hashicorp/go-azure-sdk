package comanageddevice

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

type CreateComanagedDevicePlayLostModeSoundOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateComanagedDevicePlayLostModeSoundOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateComanagedDevicePlayLostModeSoundOperationOptions() CreateComanagedDevicePlayLostModeSoundOperationOptions {
	return CreateComanagedDevicePlayLostModeSoundOperationOptions{}
}

func (o CreateComanagedDevicePlayLostModeSoundOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDevicePlayLostModeSoundOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDevicePlayLostModeSoundOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDevicePlayLostModeSound - Invoke action playLostModeSound. Play lost mode sound
func (c ComanagedDeviceClient) CreateComanagedDevicePlayLostModeSound(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input CreateComanagedDevicePlayLostModeSoundRequest, options CreateComanagedDevicePlayLostModeSoundOperationOptions) (result CreateComanagedDevicePlayLostModeSoundOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/playLostModeSound", id.ID()),
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
