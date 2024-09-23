package authenticationeventlistener

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationEventListenerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.AuthenticationEventListener
}

type CreateAuthenticationEventListenerOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationEventListenerOperationOptions() CreateAuthenticationEventListenerOperationOptions {
	return CreateAuthenticationEventListenerOperationOptions{}
}

func (o CreateAuthenticationEventListenerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationEventListenerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationEventListenerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationEventListener - Create authenticationEventListener. Create a new authenticationEventListener
// object. You can create one of the following subtypes that are derived from authenticationEventListener.
func (c AuthenticationEventListenerClient) CreateAuthenticationEventListener(ctx context.Context, input stable.AuthenticationEventListener, options CreateAuthenticationEventListenerOperationOptions) (result CreateAuthenticationEventListenerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identity/authenticationEventListeners",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalAuthenticationEventListenerImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
