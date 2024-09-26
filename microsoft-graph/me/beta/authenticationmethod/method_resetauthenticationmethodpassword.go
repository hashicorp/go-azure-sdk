package authenticationmethod

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

type ResetAuthenticationMethodPasswordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PasswordResetResponse
}

type ResetAuthenticationMethodPasswordOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetAuthenticationMethodPasswordOperationOptions() ResetAuthenticationMethodPasswordOperationOptions {
	return ResetAuthenticationMethodPasswordOperationOptions{}
}

func (o ResetAuthenticationMethodPasswordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetAuthenticationMethodPasswordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetAuthenticationMethodPasswordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetAuthenticationMethodPassword - Invoke action resetPassword. Initiate a reset for the password associated with a
// password authentication method object. This can only be done by an administrator with appropriate permissions and
// can't be performed on a user's own account. To reset a user's password in in Azure AD B2C, use the Update user API
// operation and update the passwordProfile > forceChangePasswordNextSignIn object. This flow writes the new password to
// Microsoft Entra ID and pushes it to on-premises Active Directory if configured using password writeback. The admin
// can either provide a new password or have the system generate one. The user is prompted to change their password on
// their next sign in. This reset is a long-running operation and will return a Location header with a link where the
// caller can periodically check for the status of the reset operation.
func (c AuthenticationMethodClient) ResetAuthenticationMethodPassword(ctx context.Context, id beta.MeAuthenticationMethodId, input ResetAuthenticationMethodPasswordRequest, options ResetAuthenticationMethodPasswordOperationOptions) (result ResetAuthenticationMethodPasswordOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/resetPassword", id.ID()),
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

	var model beta.PasswordResetResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
