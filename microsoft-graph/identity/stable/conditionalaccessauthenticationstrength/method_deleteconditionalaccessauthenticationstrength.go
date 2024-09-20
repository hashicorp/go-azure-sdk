package conditionalaccessauthenticationstrength

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteConditionalAccessAuthenticationStrengthOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteConditionalAccessAuthenticationStrengthOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteConditionalAccessAuthenticationStrengthOperationOptions() DeleteConditionalAccessAuthenticationStrengthOperationOptions {
	return DeleteConditionalAccessAuthenticationStrengthOperationOptions{}
}

func (o DeleteConditionalAccessAuthenticationStrengthOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteConditionalAccessAuthenticationStrengthOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteConditionalAccessAuthenticationStrengthOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteConditionalAccessAuthenticationStrength - Delete navigation property authenticationStrength for identity
func (c ConditionalAccessAuthenticationStrengthClient) DeleteConditionalAccessAuthenticationStrength(ctx context.Context, options DeleteConditionalAccessAuthenticationStrengthOperationOptions) (result DeleteConditionalAccessAuthenticationStrengthOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identity/conditionalAccess/authenticationStrength",
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
