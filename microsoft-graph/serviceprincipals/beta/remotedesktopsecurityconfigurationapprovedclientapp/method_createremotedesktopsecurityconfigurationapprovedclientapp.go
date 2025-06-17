package remotedesktopsecurityconfigurationapprovedclientapp

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

type CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ApprovedClientApp
}

type CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions() CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions {
	return CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions{}
}

func (o CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRemoteDesktopSecurityConfigurationApprovedClientApp - Create new navigation property to approvedClientApps for
// servicePrincipals
func (c RemoteDesktopSecurityConfigurationApprovedClientAppClient) CreateRemoteDesktopSecurityConfigurationApprovedClientApp(ctx context.Context, id beta.ServicePrincipalId, input beta.ApprovedClientApp, options CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationOptions) (result CreateRemoteDesktopSecurityConfigurationApprovedClientAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration/approvedClientApps", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.ApprovedClientApp
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
