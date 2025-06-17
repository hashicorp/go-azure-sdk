package customauthenticationextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationConfigurationValidation
}

type ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions() ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions {
	return ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions{}
}

func (o ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateCustomAuthenticationExtensionAuthenticationConfiguration - Invoke action validateAuthenticationConfiguration.
// An API to check validity of the endpoint and and authentication configuration for a customAuthenticationExtension
// object, which can represent one of the following derived types
func (c CustomAuthenticationExtensionClient) ValidateCustomAuthenticationExtensionAuthenticationConfiguration(ctx context.Context, id stable.IdentityCustomAuthenticationExtensionId, options ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationOptions) (result ValidateCustomAuthenticationExtensionAuthenticationConfigurationOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/validateAuthenticationConfiguration", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.AuthenticationConfigurationValidation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
