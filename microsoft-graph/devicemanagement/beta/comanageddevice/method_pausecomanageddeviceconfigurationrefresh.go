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

type PauseComanagedDeviceConfigurationRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type PauseComanagedDeviceConfigurationRefreshOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultPauseComanagedDeviceConfigurationRefreshOperationOptions() PauseComanagedDeviceConfigurationRefreshOperationOptions {
	return PauseComanagedDeviceConfigurationRefreshOperationOptions{}
}

func (o PauseComanagedDeviceConfigurationRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PauseComanagedDeviceConfigurationRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o PauseComanagedDeviceConfigurationRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PauseComanagedDeviceConfigurationRefresh - Invoke action pauseConfigurationRefresh. Initiates a command to pause
// config refresh for the device.
func (c ComanagedDeviceClient) PauseComanagedDeviceConfigurationRefresh(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input PauseComanagedDeviceConfigurationRefreshRequest, options PauseComanagedDeviceConfigurationRefreshOperationOptions) (result PauseComanagedDeviceConfigurationRefreshOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/pauseConfigurationRefresh", id.ID()),
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
