package authenticationpasswordlessmicrosoftauthenticatormethoddevice

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

type GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Device
}

type GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions() GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions {
	return GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions{}
}

func (o GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice - Get device from me
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient) GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice(ctx context.Context, id beta.MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, options GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationOptions) (result GetAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/device", id.ID()),
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

	var model beta.Device
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
