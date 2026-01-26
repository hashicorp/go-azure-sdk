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

type ComputeNodeRebootOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ComputeNodeRebootOperationOptions struct {
	ClientRequestId       *string
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultComputeNodeRebootOperationOptions() ComputeNodeRebootOperationOptions {
	return ComputeNodeRebootOperationOptions{}
}

func (o ComputeNodeRebootOperationOptions) ToHeaders() *client.Headers {
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

func (o ComputeNodeRebootOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ComputeNodeRebootOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// ComputeNodeReboot ...
func (c ComputeNodesClient) ComputeNodeReboot(ctx context.Context, id NodeId, input NodeRebootParameter, options ComputeNodeRebootOperationOptions) (result ComputeNodeRebootOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; odata=minimalmetadata; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reboot", id.Path()),
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
