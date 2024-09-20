package cloudpc

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

type ResizeCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResizeCloudPCOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultResizeCloudPCOperationOptions() ResizeCloudPCOperationOptions {
	return ResizeCloudPCOperationOptions{}
}

func (o ResizeCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResizeCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResizeCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResizeCloudPC - Invoke action resize. Upgrade or downgrade an existing Cloud PC to a configuration with a new virtual
// CPU (vCPU) and storage size.
func (c CloudPCClient) ResizeCloudPC(ctx context.Context, id beta.UserIdCloudPCId, input ResizeCloudPCRequest, options ResizeCloudPCOperationOptions) (result ResizeCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resize", id.ID()),
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
