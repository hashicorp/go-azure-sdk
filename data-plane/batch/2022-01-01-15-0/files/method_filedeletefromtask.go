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

type FileDeleteFromTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type FileDeleteFromTaskOperationOptions struct {
	ClientRequestId       *string
	OcpDate               *string
	Recursive             *bool
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultFileDeleteFromTaskOperationOptions() FileDeleteFromTaskOperationOptions {
	return FileDeleteFromTaskOperationOptions{}
}

func (o FileDeleteFromTaskOperationOptions) ToHeaders() *client.Headers {
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

func (o FileDeleteFromTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FileDeleteFromTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Recursive != nil {
		out.Append("recursive", fmt.Sprintf("%v", *o.Recursive))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// FileDeleteFromTask ...
func (c FilesClient) FileDeleteFromTask(ctx context.Context, id FileId, options FileDeleteFromTaskOperationOptions) (result FileDeleteFromTaskOperationResponse, err error) {
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
