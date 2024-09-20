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

type EnableAuthenticationPhoneMethodSmsSignInOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EnableAuthenticationPhoneMethodSmsSignInOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultEnableAuthenticationPhoneMethodSmsSignInOperationOptions() EnableAuthenticationPhoneMethodSmsSignInOperationOptions {
	return EnableAuthenticationPhoneMethodSmsSignInOperationOptions{}
}

func (o EnableAuthenticationPhoneMethodSmsSignInOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EnableAuthenticationPhoneMethodSmsSignInOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EnableAuthenticationPhoneMethodSmsSignInOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EnableAuthenticationPhoneMethodSmsSignIn - Invoke action enableSmsSignIn. Enable SMS sign-in for an existing mobile
// phone number registered to a user. To be successfully enabled
func (c AuthenticationPhoneMethodClient) EnableAuthenticationPhoneMethodSmsSignIn(ctx context.Context, id stable.MeAuthenticationPhoneMethodId, options EnableAuthenticationPhoneMethodSmsSignInOperationOptions) (result EnableAuthenticationPhoneMethodSmsSignInOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/enableSmsSignIn", id.ID()),
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
