package authenticationqrcodepinmethodstandardqrcode

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationQrCodePinMethodStandardQRCodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.QrCode
}

type GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions() GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions {
	return GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions{}
}

func (o GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationQrCodePinMethodStandardQRCode - Get standardQRCode from me. Standard QR code is primary QR code of
// the user with lifetime upto 395 days (13 months). There can be only one active standard QR code for the user.
func (c AuthenticationQrCodePinMethodStandardQRCodeClient) GetAuthenticationQrCodePinMethodStandardQRCode(ctx context.Context, options GetAuthenticationQrCodePinMethodStandardQRCodeOperationOptions) (result GetAuthenticationQrCodePinMethodStandardQRCodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/authentication/qrCodePinMethod/standardQRCode",
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

	var model beta.QrCode
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
