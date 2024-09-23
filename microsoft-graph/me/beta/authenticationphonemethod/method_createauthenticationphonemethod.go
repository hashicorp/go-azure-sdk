package authenticationphonemethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PhoneAuthenticationMethod
}

type CreateAuthenticationPhoneMethodOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationPhoneMethodOperationOptions() CreateAuthenticationPhoneMethodOperationOptions {
	return CreateAuthenticationPhoneMethodOperationOptions{}
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationPhoneMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationPhoneMethod - Create new navigation property to phoneMethods for me
func (c AuthenticationPhoneMethodClient) CreateAuthenticationPhoneMethod(ctx context.Context, input beta.PhoneAuthenticationMethod, options CreateAuthenticationPhoneMethodOperationOptions) (result CreateAuthenticationPhoneMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/authentication/phoneMethods",
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

	var model beta.PhoneAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
