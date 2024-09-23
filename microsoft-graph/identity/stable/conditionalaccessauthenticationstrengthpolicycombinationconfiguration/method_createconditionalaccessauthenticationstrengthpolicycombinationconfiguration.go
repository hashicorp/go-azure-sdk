package conditionalaccessauthenticationstrengthpolicycombinationconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.AuthenticationCombinationConfiguration
}

type CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions() CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions {
	return CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions{}
}

func (o CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfiguration - Create
// authenticationCombinationConfiguration. Create a new authenticationCombinationConfiguration object which can be of
// one of the following derived types: * fido2combinationConfiguration * x509certificatecombinationconfiguration
func (c ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient) CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfiguration(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthPolicyId, input stable.AuthenticationCombinationConfiguration, options CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationOptions) (result CreateConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/combinationConfigurations", id.ID()),
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
	model, err := stable.UnmarshalAuthenticationCombinationConfigurationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
