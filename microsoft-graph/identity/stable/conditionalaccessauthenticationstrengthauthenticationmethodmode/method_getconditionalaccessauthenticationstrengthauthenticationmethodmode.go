package conditionalaccessauthenticationstrengthauthenticationmethodmode

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationMethodModeDetail
}

type GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions() GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions {
	return GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions{}
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConditionalAccessAuthenticationStrengthAuthenticationMethodMode - Get authenticationMethodModes from identity.
// Names and descriptions of all valid authentication method modes in the system.
func (c ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient) GetConditionalAccessAuthenticationStrengthAuthenticationMethodMode(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthAuthenticationMethodModeId, options GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationOptions) (result GetConditionalAccessAuthenticationStrengthAuthenticationMethodModeOperationResponse, err error) {
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

	var model stable.AuthenticationMethodModeDetail
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
