package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileDeleteFromComputeNodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type FileDeleteFromComputeNodeOperationOptions struct {
	ClientRequestId       *string
	OcpDate               *string
	Recursive             *bool
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultFileDeleteFromComputeNodeOperationOptions() FileDeleteFromComputeNodeOperationOptions {
	return FileDeleteFromComputeNodeOperationOptions{}
}

func (o FileDeleteFromComputeNodeOperationOptions) ToHeaders() *client.Headers {
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

func (o FileDeleteFromComputeNodeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FileDeleteFromComputeNodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Recursive != nil {
		out.Append("recursive", fmt.Sprintf("%v", *o.Recursive))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// FileDeleteFromComputeNode ...
func (c FilesClient) FileDeleteFromComputeNode(ctx context.Context, id NodeFileId, options FileDeleteFromComputeNodeOperationOptions) (result FileDeleteFromComputeNodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; odata=minimalmetadata; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.Path(),
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
