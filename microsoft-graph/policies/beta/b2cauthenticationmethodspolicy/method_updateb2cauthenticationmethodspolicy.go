package b2cauthenticationmethodspolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateB2cAuthenticationMethodsPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2cAuthenticationMethodsPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateB2cAuthenticationMethodsPolicyOperationOptions() UpdateB2cAuthenticationMethodsPolicyOperationOptions {
	return UpdateB2cAuthenticationMethodsPolicyOperationOptions{}
}

func (o UpdateB2cAuthenticationMethodsPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2cAuthenticationMethodsPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2cAuthenticationMethodsPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2cAuthenticationMethodsPolicy - Update b2cAuthenticationMethodsPolicy. Update the properties of a
// b2cAuthenticationMethodsPolicy object.
func (c B2cAuthenticationMethodsPolicyClient) UpdateB2cAuthenticationMethodsPolicy(ctx context.Context, input beta.B2cAuthenticationMethodsPolicy, options UpdateB2cAuthenticationMethodsPolicyOperationOptions) (result UpdateB2cAuthenticationMethodsPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/b2cAuthenticationMethodsPolicy",
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
