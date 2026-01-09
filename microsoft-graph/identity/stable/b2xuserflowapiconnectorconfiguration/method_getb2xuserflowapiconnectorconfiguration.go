package b2xuserflowapiconnectorconfiguration

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

type GetB2xUserFlowApiConnectorConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserFlowApiConnectorConfiguration
}

type GetB2xUserFlowApiConnectorConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetB2xUserFlowApiConnectorConfigurationOperationOptions() GetB2xUserFlowApiConnectorConfigurationOperationOptions {
	return GetB2xUserFlowApiConnectorConfigurationOperationOptions{}
}

func (o GetB2xUserFlowApiConnectorConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowApiConnectorConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetB2xUserFlowApiConnectorConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowApiConnectorConfiguration - Get userFlowApiConnectorConfiguration. Get the apiConnectorConfiguration
// property in a b2xIdentityUserFlow to detail the API connectors enabled for the user flow.
func (c B2xUserFlowApiConnectorConfigurationClient) GetB2xUserFlowApiConnectorConfiguration(ctx context.Context, id stable.IdentityB2xUserFlowId, options GetB2xUserFlowApiConnectorConfigurationOperationOptions) (result GetB2xUserFlowApiConnectorConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration", id.ID()),
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

	var model stable.UserFlowApiConnectorConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
