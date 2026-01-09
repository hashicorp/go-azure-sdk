package authenticationmethoddevicehardwareoathdevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationMethodDeviceHardwareOathDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HardwareOathTokenAuthenticationMethodDevice
}

type CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions() CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions {
	return CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions{}
}

func (o CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationMethodDeviceHardwareOathDevice - Create hardwareOathTokenAuthenticationMethodDevice. Create a new
// hardwareOathTokenAuthenticationMethodDevice object. You can optionally create and assign to a user in the same
// request; Or assign to a user via the assign API.
func (c AuthenticationMethodDeviceHardwareOathDeviceClient) CreateAuthenticationMethodDeviceHardwareOathDevice(ctx context.Context, input beta.HardwareOathTokenAuthenticationMethodDevice, options CreateAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) (result CreateAuthenticationMethodDeviceHardwareOathDeviceOperationResponse, err error) {
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
		Path:          "/directory/authenticationMethodDevices/hardwareOathDevices",
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

	var model beta.HardwareOathTokenAuthenticationMethodDevice
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
