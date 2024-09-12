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

type DeleteMicrosoftTunnelSiteConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteMicrosoftTunnelSiteConfigurationOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteMicrosoftTunnelSiteConfigurationOperationOptions() DeleteMicrosoftTunnelSiteConfigurationOperationOptions {
	return DeleteMicrosoftTunnelSiteConfigurationOperationOptions{}
}

func (o DeleteMicrosoftTunnelSiteConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteMicrosoftTunnelSiteConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteMicrosoftTunnelSiteConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteMicrosoftTunnelSiteConfiguration - Delete navigation property microsoftTunnelConfiguration for deviceManagement
func (c MicrosoftTunnelSiteMicrosoftTunnelConfigurationClient) DeleteMicrosoftTunnelSiteConfiguration(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteId, options DeleteMicrosoftTunnelSiteConfigurationOperationOptions) (result DeleteMicrosoftTunnelSiteConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/microsoftTunnelConfiguration", id.ID()),
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
