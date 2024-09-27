package authenticationphonemethod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisableAuthenticationPhoneMethodSmsSignInOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DisableAuthenticationPhoneMethodSmsSignInOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDisableAuthenticationPhoneMethodSmsSignInOperationOptions() DisableAuthenticationPhoneMethodSmsSignInOperationOptions {
	return DisableAuthenticationPhoneMethodSmsSignInOperationOptions{}
}

func (o DisableAuthenticationPhoneMethodSmsSignInOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DisableAuthenticationPhoneMethodSmsSignInOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DisableAuthenticationPhoneMethodSmsSignInOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DisableAuthenticationPhoneMethodSmsSignIn - Invoke action disableSmsSignIn. Disable SMS sign-in for an existing
// mobile phone number registered to a user. The number will no longer be available for SMS sign-in, which can prevent
// your user from signing in.
func (c AuthenticationPhoneMethodClient) DisableAuthenticationPhoneMethodSmsSignIn(ctx context.Context, id stable.UserIdAuthenticationPhoneMethodId, options DisableAuthenticationPhoneMethodSmsSignInOperationOptions) (result DisableAuthenticationPhoneMethodSmsSignInOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/disableSmsSignIn", id.ID()),
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

	return
}
