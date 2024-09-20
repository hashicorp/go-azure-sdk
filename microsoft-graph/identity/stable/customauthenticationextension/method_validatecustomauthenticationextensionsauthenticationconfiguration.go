package customauthenticationextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationConfigurationValidation
}

type ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions() ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions {
	return ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions{}
}

func (o ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateCustomAuthenticationExtensionsAuthenticationConfiguration - Invoke action validateAuthenticationConfiguration
func (c CustomAuthenticationExtensionClient) ValidateCustomAuthenticationExtensionsAuthenticationConfiguration(ctx context.Context, input ValidateCustomAuthenticationExtensionsAuthenticationConfigurationRequest, options ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationOptions) (result ValidateCustomAuthenticationExtensionsAuthenticationConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identity/customAuthenticationExtensions/validateAuthenticationConfiguration",
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

	var model stable.AuthenticationConfigurationValidation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
