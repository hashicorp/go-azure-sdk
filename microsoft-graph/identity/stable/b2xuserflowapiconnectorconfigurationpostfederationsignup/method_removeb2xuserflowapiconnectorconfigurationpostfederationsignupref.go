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

type RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions() RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions {
	return RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions{}
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef - Delete ref of navigation property
// postFederationSignup for identity
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx context.Context, id stable.IdentityB2xUserFlowId, options RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions) (result RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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
