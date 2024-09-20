package authenticationmethodspolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationMethodsPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationMethodsPolicyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAuthenticationMethodsPolicyOperationOptions() UpdateAuthenticationMethodsPolicyOperationOptions {
	return UpdateAuthenticationMethodsPolicyOperationOptions{}
}

func (o UpdateAuthenticationMethodsPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationMethodsPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationMethodsPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationMethodsPolicy - Update authenticationMethodsPolicy. Update the properties of an
// authenticationMethodsPolicy object.
func (c AuthenticationMethodsPolicyClient) UpdateAuthenticationMethodsPolicy(ctx context.Context, input stable.AuthenticationMethodsPolicy, options UpdateAuthenticationMethodsPolicyOperationOptions) (result UpdateAuthenticationMethodsPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/authenticationMethodsPolicy",
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
