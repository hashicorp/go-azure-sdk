package authenticationmethodspolicyauthenticationmethodconfiguration

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationMethodsPolicyConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.AuthenticationMethodConfiguration
}

type CreateAuthenticationMethodsPolicyConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationMethodsPolicyConfigurationOperationOptions() CreateAuthenticationMethodsPolicyConfigurationOperationOptions {
	return CreateAuthenticationMethodsPolicyConfigurationOperationOptions{}
}

func (o CreateAuthenticationMethodsPolicyConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationMethodsPolicyConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationMethodsPolicyConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationMethodsPolicyConfiguration - Create new navigation property to authenticationMethodConfigurations
// for policies
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) CreateAuthenticationMethodsPolicyConfiguration(ctx context.Context, input beta.AuthenticationMethodConfiguration, options CreateAuthenticationMethodsPolicyConfigurationOperationOptions) (result CreateAuthenticationMethodsPolicyConfigurationOperationResponse, err error) {
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
		Path:          "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalAuthenticationMethodConfigurationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
