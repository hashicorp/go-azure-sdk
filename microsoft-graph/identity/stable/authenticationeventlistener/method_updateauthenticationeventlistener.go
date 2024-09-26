package authenticationeventlistener

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationEventListenerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationEventListenerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationEventListenerOperationOptions() UpdateAuthenticationEventListenerOperationOptions {
	return UpdateAuthenticationEventListenerOperationOptions{}
}

func (o UpdateAuthenticationEventListenerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationEventListenerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationEventListenerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationEventListener - Update authenticationEventListener. Update the properties of an
// authenticationEventListener object. You must specify the @odata.type property and the value of the
// authenticationEventListener object type to update.
func (c AuthenticationEventListenerClient) UpdateAuthenticationEventListener(ctx context.Context, id stable.IdentityAuthenticationEventListenerId, input stable.AuthenticationEventListener, options UpdateAuthenticationEventListenerOperationOptions) (result UpdateAuthenticationEventListenerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
