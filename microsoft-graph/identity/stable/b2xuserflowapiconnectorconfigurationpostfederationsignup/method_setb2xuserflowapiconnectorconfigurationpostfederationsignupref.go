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

type SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions() SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions {
	return SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions{}
}

func (o SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef - Update the ref of navigation property
// postFederationSignup in identity
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.ReferenceUpdate, options SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) (result SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup/$ref", id.ID()),
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
