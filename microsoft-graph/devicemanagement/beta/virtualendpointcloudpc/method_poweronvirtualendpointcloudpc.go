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

type PowerOnVirtualEndpointCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type PowerOnVirtualEndpointCloudPCOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultPowerOnVirtualEndpointCloudPCOperationOptions() PowerOnVirtualEndpointCloudPCOperationOptions {
	return PowerOnVirtualEndpointCloudPCOperationOptions{}
}

func (o PowerOnVirtualEndpointCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PowerOnVirtualEndpointCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o PowerOnVirtualEndpointCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PowerOnVirtualEndpointCloudPC - Invoke action powerOn. Power on a Windows 365 Frontline Cloud PC. This action
// supports Microsoft Endpoint Manager (MEM) admin scenarios. After a Windows 365 Frontline Cloud PC is powered on, it
// is allocated to a user, and licenses are assigned immediately. Only IT admin users can perform this action.
func (c VirtualEndpointCloudPCClient) PowerOnVirtualEndpointCloudPC(ctx context.Context, id beta.DeviceManagementVirtualEndpointCloudPCId, options PowerOnVirtualEndpointCloudPCOperationOptions) (result PowerOnVirtualEndpointCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/powerOn", id.ID()),
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
