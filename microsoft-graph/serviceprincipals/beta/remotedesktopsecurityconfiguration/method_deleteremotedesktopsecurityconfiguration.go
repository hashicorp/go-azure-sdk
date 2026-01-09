package remotedesktopsecurityconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteRemoteDesktopSecurityConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteRemoteDesktopSecurityConfigurationOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteRemoteDesktopSecurityConfigurationOperationOptions() DeleteRemoteDesktopSecurityConfigurationOperationOptions {
	return DeleteRemoteDesktopSecurityConfigurationOperationOptions{}
}

func (o DeleteRemoteDesktopSecurityConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteRemoteDesktopSecurityConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteRemoteDesktopSecurityConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteRemoteDesktopSecurityConfiguration - Delete remoteDesktopSecurityConfiguration. Delete a
// remoteDesktopSecurityConfiguration object on a servicePrincipal. Removing remoteDesktopSecurityConfiguration object
// on the servicePrincipal disables the Microsoft Entra ID Remote Desktop Services (RDS) authentication protocol to
// authenticate a user to Microsoft Entra joined or Microsoft Entra hybrid joined devices, and removes any target device
// groups that you configured for SSO.
func (c RemoteDesktopSecurityConfigurationClient) DeleteRemoteDesktopSecurityConfiguration(ctx context.Context, id beta.ServicePrincipalId, options DeleteRemoteDesktopSecurityConfigurationOperationOptions) (result DeleteRemoteDesktopSecurityConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
