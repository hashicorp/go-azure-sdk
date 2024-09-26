package remotedesktopsecurityconfigurationtargetdevicegroup

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

type DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions() DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions {
	return DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions{}
}

func (o DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroup - Delete targetDeviceGroup. Delete a targetDeviceGroup
// object for the remoteDesktopSecurityConfiguration object on the servicePrincipal. Any user authenticating using the
// Microsoft Entra ID Remote Desktop Services (RDS) authentication protocol to a Microsoft Entra joined or Microsoft
// Entra hybrid joined device that's in the removed targetDeviceGroup doesn't get SSO prompts.
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroup(ctx context.Context, id beta.ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId, options DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) (result DeleteRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
