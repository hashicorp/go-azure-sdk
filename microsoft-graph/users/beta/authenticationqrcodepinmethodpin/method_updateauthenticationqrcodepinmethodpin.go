package authenticationqrcodepinmethodpin

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

type UpdateAuthenticationQrCodePinMethodPinOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationQrCodePinMethodPinOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationQrCodePinMethodPinOperationOptions() UpdateAuthenticationQrCodePinMethodPinOperationOptions {
	return UpdateAuthenticationQrCodePinMethodPinOperationOptions{}
}

func (o UpdateAuthenticationQrCodePinMethodPinOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationQrCodePinMethodPinOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationQrCodePinMethodPinOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationQrCodePinMethodPin - Update qrPin. Reset a user's PIN and generate a new temporary PIN that's
// represented by a qrPin object and is linked to the user's QR Code authentication method object.
func (c AuthenticationQrCodePinMethodPinClient) UpdateAuthenticationQrCodePinMethodPin(ctx context.Context, id beta.UserId, input beta.QrPin, options UpdateAuthenticationQrCodePinMethodPinOperationOptions) (result UpdateAuthenticationQrCodePinMethodPinOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/authentication/qrCodePinMethod/pin", id.ID()),
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
