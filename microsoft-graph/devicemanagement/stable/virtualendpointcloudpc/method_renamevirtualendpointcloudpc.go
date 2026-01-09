package virtualendpointcloudpc

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

type RenameVirtualEndpointCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RenameVirtualEndpointCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRenameVirtualEndpointCloudPCOperationOptions() RenameVirtualEndpointCloudPCOperationOptions {
	return RenameVirtualEndpointCloudPCOperationOptions{}
}

func (o RenameVirtualEndpointCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RenameVirtualEndpointCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RenameVirtualEndpointCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RenameVirtualEndpointCloudPC - Invoke action rename. Rename a specific cloudPC object. Use this API to update the
// displayName of a Cloud PC entity.
func (c VirtualEndpointCloudPCClient) RenameVirtualEndpointCloudPC(ctx context.Context, id stable.DeviceManagementVirtualEndpointCloudPCId, input RenameVirtualEndpointCloudPCRequest, options RenameVirtualEndpointCloudPCOperationOptions) (result RenameVirtualEndpointCloudPCOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/rename", id.ID()),
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

	return
}
