package computenodes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeGetRemoteDesktopOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type ComputeNodeGetRemoteDesktopOperationOptions struct {
	ClientRequestId       *string
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultComputeNodeGetRemoteDesktopOperationOptions() ComputeNodeGetRemoteDesktopOperationOptions {
	return ComputeNodeGetRemoteDesktopOperationOptions{}
}

func (o ComputeNodeGetRemoteDesktopOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o ComputeNodeGetRemoteDesktopOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ComputeNodeGetRemoteDesktopOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// ComputeNodeGetRemoteDesktop ...
func (c ComputeNodesClient) ComputeNodeGetRemoteDesktop(ctx context.Context, id NodeId, options ComputeNodeGetRemoteDesktopOperationOptions) (result ComputeNodeGetRemoteDesktopOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/rdp", id.Path()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
