package authenticationqrcodepinmethodpin

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationQrCodePinMethodPinPinOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationQrCodePinMethodPinPinOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationQrCodePinMethodPinPinOperationOptions() UpdateAuthenticationQrCodePinMethodPinPinOperationOptions {
	return UpdateAuthenticationQrCodePinMethodPinPinOperationOptions{}
}

func (o UpdateAuthenticationQrCodePinMethodPinPinOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationQrCodePinMethodPinPinOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationQrCodePinMethodPinPinOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationQrCodePinMethodPinPin - Invoke action updatePin
func (c AuthenticationQrCodePinMethodPinClient) UpdateAuthenticationQrCodePinMethodPinPin(ctx context.Context, input UpdateAuthenticationQrCodePinMethodPinPinRequest, options UpdateAuthenticationQrCodePinMethodPinPinOperationOptions) (result UpdateAuthenticationQrCodePinMethodPinPinOperationResponse, err error) {
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
		Path:          "/me/authentication/qrCodePinMethod/pin/updatePin",
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

	return
}
