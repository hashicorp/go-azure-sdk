package conditionalaccessauthenticationstrength

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateConditionalAccessAuthenticationStrengthOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateConditionalAccessAuthenticationStrengthOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateConditionalAccessAuthenticationStrengthOperationOptions() UpdateConditionalAccessAuthenticationStrengthOperationOptions {
	return UpdateConditionalAccessAuthenticationStrengthOperationOptions{}
}

func (o UpdateConditionalAccessAuthenticationStrengthOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateConditionalAccessAuthenticationStrengthOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateConditionalAccessAuthenticationStrengthOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateConditionalAccessAuthenticationStrength - Update the navigation property authenticationStrength in identity
func (c ConditionalAccessAuthenticationStrengthClient) UpdateConditionalAccessAuthenticationStrength(ctx context.Context, input stable.AuthenticationStrengthRoot, options UpdateConditionalAccessAuthenticationStrengthOperationOptions) (result UpdateConditionalAccessAuthenticationStrengthOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/identity/conditionalAccess/authenticationStrength",
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
