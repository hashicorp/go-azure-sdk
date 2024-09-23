package microsofttunnelconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMicrosoftTunnelConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MicrosoftTunnelConfiguration
}

type GetMicrosoftTunnelConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetMicrosoftTunnelConfigurationOperationOptions() GetMicrosoftTunnelConfigurationOperationOptions {
	return GetMicrosoftTunnelConfigurationOperationOptions{}
}

func (o GetMicrosoftTunnelConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMicrosoftTunnelConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetMicrosoftTunnelConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMicrosoftTunnelConfiguration - Get microsoftTunnelConfigurations from deviceManagement. Collection of
// MicrosoftTunnelConfiguration settings associated with account.
func (c MicrosoftTunnelConfigurationClient) GetMicrosoftTunnelConfiguration(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelConfigurationId, options GetMicrosoftTunnelConfigurationOperationOptions) (result GetMicrosoftTunnelConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.MicrosoftTunnelConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
