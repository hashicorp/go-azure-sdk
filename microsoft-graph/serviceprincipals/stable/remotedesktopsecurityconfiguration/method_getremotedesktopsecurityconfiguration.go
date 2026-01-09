package remotedesktopsecurityconfiguration

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

type GetRemoteDesktopSecurityConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.RemoteDesktopSecurityConfiguration
}

type GetRemoteDesktopSecurityConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetRemoteDesktopSecurityConfigurationOperationOptions() GetRemoteDesktopSecurityConfigurationOperationOptions {
	return GetRemoteDesktopSecurityConfigurationOperationOptions{}
}

func (o GetRemoteDesktopSecurityConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRemoteDesktopSecurityConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetRemoteDesktopSecurityConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRemoteDesktopSecurityConfiguration - Get remoteDesktopSecurityConfiguration. Read the properties and relationships
// of a remoteDesktopSecurityConfiguration object on a servicePrincipal. Use this configuration to view the Microsoft
// Entra ID Remote Desktop Services (RDS) authentication protocol to authenticate a user to Microsoft Entra joined or
// Microsoft Entra hybrid joined devices. Additionally you can view any targetDeviceGroups that have been configured for
// SSO.
func (c RemoteDesktopSecurityConfigurationClient) GetRemoteDesktopSecurityConfiguration(ctx context.Context, id stable.ServicePrincipalId, options GetRemoteDesktopSecurityConfigurationOperationOptions) (result GetRemoteDesktopSecurityConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration", id.ID()),
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

	var model stable.RemoteDesktopSecurityConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
