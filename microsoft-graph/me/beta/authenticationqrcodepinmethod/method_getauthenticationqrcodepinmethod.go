package authenticationqrcodepinmethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationQrCodePinMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.QrCodePinAuthenticationMethod
}

type GetAuthenticationQrCodePinMethodOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationQrCodePinMethodOperationOptions() GetAuthenticationQrCodePinMethodOperationOptions {
	return GetAuthenticationQrCodePinMethodOperationOptions{}
}

func (o GetAuthenticationQrCodePinMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationQrCodePinMethodOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationQrCodePinMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationQrCodePinMethod - Get qrCodePinMethod from me. Represents a QR code authentication method registered
// to a user for authentication.
func (c AuthenticationQrCodePinMethodClient) GetAuthenticationQrCodePinMethod(ctx context.Context, options GetAuthenticationQrCodePinMethodOperationOptions) (result GetAuthenticationQrCodePinMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/authentication/qrCodePinMethod",
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

	var model beta.QrCodePinAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
