package authenticationsigninpreference

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationSignInPreferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SignInPreferences
}

type GetAuthenticationSignInPreferenceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAuthenticationSignInPreferenceOperationOptions() GetAuthenticationSignInPreferenceOperationOptions {
	return GetAuthenticationSignInPreferenceOperationOptions{}
}

func (o GetAuthenticationSignInPreferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationSignInPreferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAuthenticationSignInPreferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationSignInPreference - Get signInPreferences property value. The settings and preferences for the
// sign-in experience of a user. Use this property to configure the user's default multifactor authentication (MFA)
// method.
func (c AuthenticationSignInPreferenceClient) GetAuthenticationSignInPreference(ctx context.Context, options GetAuthenticationSignInPreferenceOperationOptions) (result GetAuthenticationSignInPreferenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/authentication/signInPreferences",
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

	var model beta.SignInPreferences
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
