package authenticationmethodspolicyauthenticationmethodconfiguration

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationMethodsPolicyConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.AuthenticationMethodConfiguration
}

type GetAuthenticationMethodsPolicyConfigurationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetAuthenticationMethodsPolicyConfigurationOperationOptions() GetAuthenticationMethodsPolicyConfigurationOperationOptions {
	return GetAuthenticationMethodsPolicyConfigurationOperationOptions{}
}

func (o GetAuthenticationMethodsPolicyConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationMethodsPolicyConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationMethodsPolicyConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationMethodsPolicyConfiguration - Get authenticationMethodConfigurations from policies. Represents the
// settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) GetAuthenticationMethodsPolicyConfiguration(ctx context.Context, id stable.PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationId, options GetAuthenticationMethodsPolicyConfigurationOperationOptions) (result GetAuthenticationMethodsPolicyConfigurationOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalAuthenticationMethodConfigurationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
