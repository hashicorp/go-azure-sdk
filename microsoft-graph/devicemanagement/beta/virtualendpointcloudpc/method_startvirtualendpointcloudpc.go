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

type StartVirtualEndpointCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StartVirtualEndpointCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStartVirtualEndpointCloudPCOperationOptions() StartVirtualEndpointCloudPCOperationOptions {
	return StartVirtualEndpointCloudPCOperationOptions{}
}

func (o StartVirtualEndpointCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartVirtualEndpointCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StartVirtualEndpointCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StartVirtualEndpointCloudPC - Invoke action start. Start a specific Cloud PC for a user. Currently, only Windows 365
// Frontline Cloud PCs are supported.
func (c VirtualEndpointCloudPCClient) StartVirtualEndpointCloudPC(ctx context.Context, id beta.DeviceManagementVirtualEndpointCloudPCId, options StartVirtualEndpointCloudPCOperationOptions) (result StartVirtualEndpointCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/start", id.ID()),
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
