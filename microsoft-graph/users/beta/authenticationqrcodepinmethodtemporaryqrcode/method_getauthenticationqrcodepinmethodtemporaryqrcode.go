package authenticationqrcodepinmethodtemporaryqrcode

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

type GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.QrCode
}

type GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions() GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions {
	return GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions{}
}

func (o GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationQrCodePinMethodTemporaryQRCode - Get temporaryQRCode from users. Temporary QR code has lifetime up
// to 12 hours. It can be issued when the user doesn't have access to their standard QR code. There can be only one
// active temporary QR code for the user.
func (c AuthenticationQrCodePinMethodTemporaryQRCodeClient) GetAuthenticationQrCodePinMethodTemporaryQRCode(ctx context.Context, id beta.UserId, options GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationOptions) (result GetAuthenticationQrCodePinMethodTemporaryQRCodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/authentication/qrCodePinMethod/temporaryQRCode", id.ID()),
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
