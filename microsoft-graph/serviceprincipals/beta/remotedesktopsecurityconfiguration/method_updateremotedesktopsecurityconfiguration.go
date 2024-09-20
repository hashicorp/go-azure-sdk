package remotedesktopsecurityconfiguration

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

type UpdateRemoteDesktopSecurityConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateRemoteDesktopSecurityConfigurationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateRemoteDesktopSecurityConfigurationOperationOptions() UpdateRemoteDesktopSecurityConfigurationOperationOptions {
	return UpdateRemoteDesktopSecurityConfigurationOperationOptions{}
}

func (o UpdateRemoteDesktopSecurityConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateRemoteDesktopSecurityConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateRemoteDesktopSecurityConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateRemoteDesktopSecurityConfiguration - Update remoteDesktopSecurityConfiguration. Update the properties of a
// remoteDesktopSecurityConfiguration object on the servicePrincipal. Use this configuration to enable or disable the
// Microsoft Entra ID Remote Desktop Services (RDS) authentication protocol to authenticate a user to Microsoft Entra
// joined or Microsoft Entra hybrid joined devices.
func (c RemoteDesktopSecurityConfigurationClient) UpdateRemoteDesktopSecurityConfiguration(ctx context.Context, id beta.ServicePrincipalId, input beta.RemoteDesktopSecurityConfiguration, options UpdateRemoteDesktopSecurityConfigurationOperationOptions) (result UpdateRemoteDesktopSecurityConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration", id.ID()),
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
