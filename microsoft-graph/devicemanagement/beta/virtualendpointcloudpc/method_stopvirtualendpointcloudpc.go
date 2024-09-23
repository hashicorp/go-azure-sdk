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

type StopVirtualEndpointCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopVirtualEndpointCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStopVirtualEndpointCloudPCOperationOptions() StopVirtualEndpointCloudPCOperationOptions {
	return StopVirtualEndpointCloudPCOperationOptions{}
}

func (o StopVirtualEndpointCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopVirtualEndpointCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopVirtualEndpointCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopVirtualEndpointCloudPC - Invoke action stop. Stop a specific Cloud PC for a user. Currently, only Windows 365
// Frontline Cloud PCs are supported.
func (c VirtualEndpointCloudPCClient) StopVirtualEndpointCloudPC(ctx context.Context, id beta.DeviceManagementVirtualEndpointCloudPCId, options StopVirtualEndpointCloudPCOperationOptions) (result StopVirtualEndpointCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/stop", id.ID()),
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
