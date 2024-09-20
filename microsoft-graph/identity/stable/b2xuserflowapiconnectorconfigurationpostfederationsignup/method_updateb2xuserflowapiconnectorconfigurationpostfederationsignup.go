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

type UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions() UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions {
	return UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions{}
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignup - Update the navigation property postFederationSignup
// in identity
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignup(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.IdentityApiConnector, options UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) (result UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup", id.ID()),
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
