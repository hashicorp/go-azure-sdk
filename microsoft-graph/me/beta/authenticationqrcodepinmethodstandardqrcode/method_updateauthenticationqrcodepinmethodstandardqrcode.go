package authenticationqrcodepinmethodstandardqrcode

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions() UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions {
	return UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions{}
}

func (o UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationQrCodePinMethodStandardQRCode - Create or Update QR code. Create a standard or temporary QR code,
// if there is no active QR code, or update a standard QR code. Only the expireDateTime property can be updated for a
// standard QR code.
func (c AuthenticationQrCodePinMethodStandardQRCodeClient) UpdateAuthenticationQrCodePinMethodStandardQRCode(ctx context.Context, input beta.QrCode, options UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) (result UpdateAuthenticationQrCodePinMethodStandardQRCodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/me/authentication/qrCodePinMethod/standardQRCode",
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
