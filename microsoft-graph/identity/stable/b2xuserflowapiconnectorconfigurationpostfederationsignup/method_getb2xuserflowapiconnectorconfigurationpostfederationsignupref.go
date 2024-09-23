package b2xuserflowapiconnectorconfigurationpostfederationsignup

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

type GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions() GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions {
	return GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions{}
}

func (o GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef - Get ref of postFederationSignup from identity
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx context.Context, id stable.IdentityB2xUserFlowId, options GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) (result GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup/$ref", id.ID()),
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

	return
}
