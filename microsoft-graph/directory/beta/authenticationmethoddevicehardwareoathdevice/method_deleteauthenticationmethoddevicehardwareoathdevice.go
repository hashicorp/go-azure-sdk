package authenticationmethoddevicehardwareoathdevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions() DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions {
	return DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions{}
}

func (o DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAuthenticationMethodDeviceHardwareOathDevice - Delete hardwareOathTokenAuthenticationMethodDevice. Delete a
// Hardware OATH token. Token needs to be unassigned.
func (c AuthenticationMethodDeviceHardwareOathDeviceClient) DeleteAuthenticationMethodDeviceHardwareOathDevice(ctx context.Context, id beta.DirectoryAuthenticationMethodDeviceHardwareOathDeviceId, options DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationOptions) (result DeleteAuthenticationMethodDeviceHardwareOathDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
