package b2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions() DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions {
	return DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions{}
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignup - Delete navigation property postFederationSignup for
// identity
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignup(ctx context.Context, id stable.IdentityB2xUserFlowId, options DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions) (result DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup", id.ID()),
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
