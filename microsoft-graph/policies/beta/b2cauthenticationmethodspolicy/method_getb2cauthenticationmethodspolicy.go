package b2cauthenticationmethodspolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2cAuthenticationMethodsPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.B2cAuthenticationMethodsPolicy
}

type GetB2cAuthenticationMethodsPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetB2cAuthenticationMethodsPolicyOperationOptions() GetB2cAuthenticationMethodsPolicyOperationOptions {
	return GetB2cAuthenticationMethodsPolicyOperationOptions{}
}

func (o GetB2cAuthenticationMethodsPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2cAuthenticationMethodsPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetB2cAuthenticationMethodsPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2cAuthenticationMethodsPolicy - Get b2cAuthenticationMethodsPolicy. Read the properties of a
// b2cAuthenticationMethodsPolicy object.
func (c B2cAuthenticationMethodsPolicyClient) GetB2cAuthenticationMethodsPolicy(ctx context.Context, options GetB2cAuthenticationMethodsPolicyOperationOptions) (result GetB2cAuthenticationMethodsPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/policies/b2cAuthenticationMethodsPolicy",
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

	var model beta.B2cAuthenticationMethodsPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
