package microsofttunnelsitemicrosofttunnelconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMicrosoftTunnelSiteConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMicrosoftTunnelSiteConfigurationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateMicrosoftTunnelSiteConfigurationOperationOptions() UpdateMicrosoftTunnelSiteConfigurationOperationOptions {
	return UpdateMicrosoftTunnelSiteConfigurationOperationOptions{}
}

func (o UpdateMicrosoftTunnelSiteConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMicrosoftTunnelSiteConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMicrosoftTunnelSiteConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMicrosoftTunnelSiteConfiguration - Update the navigation property microsoftTunnelConfiguration in
// deviceManagement
func (c MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient) UpdateMicrosoftTunnelSiteConfiguration(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteId, input beta.MicrosoftTunnelConfiguration, options UpdateMicrosoftTunnelSiteConfigurationOperationOptions) (result UpdateMicrosoftTunnelSiteConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/microsoftTunnelConfiguration", id.ID()),
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
