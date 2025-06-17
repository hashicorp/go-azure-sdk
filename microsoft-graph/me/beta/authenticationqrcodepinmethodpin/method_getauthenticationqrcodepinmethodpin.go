package authenticationqrcodepinmethodpin

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationQrCodePinMethodPinOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.QrPin
}

type GetAuthenticationQrCodePinMethodPinOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationQrCodePinMethodPinOperationOptions() GetAuthenticationQrCodePinMethodPinOperationOptions {
	return GetAuthenticationQrCodePinMethodPinOperationOptions{}
}

func (o GetAuthenticationQrCodePinMethodPinOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationQrCodePinMethodPinOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationQrCodePinMethodPinOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationQrCodePinMethodPin - Get pin from me. The PIN linked to the QR Code auth method of the user.
func (c AuthenticationQrCodePinMethodPinClient) GetAuthenticationQrCodePinMethodPin(ctx context.Context, options GetAuthenticationQrCodePinMethodPinOperationOptions) (result GetAuthenticationQrCodePinMethodPinOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/authentication/qrCodePinMethod/pin",
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

	var model beta.QrPin
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
