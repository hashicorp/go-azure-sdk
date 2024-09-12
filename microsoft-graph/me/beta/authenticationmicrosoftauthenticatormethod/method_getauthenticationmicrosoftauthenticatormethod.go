package authenticationmicrosoftauthenticatormethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationMicrosoftAuthenticatorMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MicrosoftAuthenticatorAuthenticationMethod
}

type GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAuthenticationMicrosoftAuthenticatorMethodOperationOptions() GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions {
	return GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions{}
}

func (o GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationMicrosoftAuthenticatorMethod - Get microsoftAuthenticatorMethods from me. The details of the
// Microsoft Authenticator app registered to a user for authentication.
func (c AuthenticationMicrosoftAuthenticatorMethodClient) GetAuthenticationMicrosoftAuthenticatorMethod(ctx context.Context, id beta.MeAuthenticationMicrosoftAuthenticatorMethodId, options GetAuthenticationMicrosoftAuthenticatorMethodOperationOptions) (result GetAuthenticationMicrosoftAuthenticatorMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.MicrosoftAuthenticatorAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
