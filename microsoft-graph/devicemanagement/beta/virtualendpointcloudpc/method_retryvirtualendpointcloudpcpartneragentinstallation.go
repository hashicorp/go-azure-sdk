package virtualendpointcloudpc

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

type RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions() RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions {
	return RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions{}
}

func (o RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RetryVirtualEndpointCloudPCPartnerAgentInstallation - Invoke action retryPartnerAgentInstallation. Retry installation
// for the partner agents that failed to install on the Cloud PC. Service side checks which agent installation failed
// firstly and retry.
func (c VirtualEndpointCloudPCClient) RetryVirtualEndpointCloudPCPartnerAgentInstallation(ctx context.Context, id beta.DeviceManagementVirtualEndpointCloudPCId, options RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationOptions) (result RetryVirtualEndpointCloudPCPartnerAgentInstallationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/retryPartnerAgentInstallation", id.ID()),
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
