package authenticationsigninpreference

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

type UpdateAuthenticationSignInPreferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationSignInPreferenceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAuthenticationSignInPreferenceOperationOptions() UpdateAuthenticationSignInPreferenceOperationOptions {
	return UpdateAuthenticationSignInPreferenceOperationOptions{}
}

func (o UpdateAuthenticationSignInPreferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationSignInPreferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationSignInPreferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationSignInPreference - Update authentication method states. Update the properties of a user's
// authentication method states. Use this API to update the following information
func (c AuthenticationSignInPreferenceClient) UpdateAuthenticationSignInPreference(ctx context.Context, id beta.UserId, input beta.SignInPreferences, options UpdateAuthenticationSignInPreferenceOperationOptions) (result UpdateAuthenticationSignInPreferenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/authentication/signInPreferences", id.ID()),
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
