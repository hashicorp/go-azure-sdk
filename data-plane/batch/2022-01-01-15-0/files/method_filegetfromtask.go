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

type FileGetFromTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type FileGetFromTaskOperationOptions struct {
	ClientRequestId       *string
	IfModifiedSince       *string
	IfUnmodifiedSince     *string
	OcpDate               *string
	OcpRange              *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultFileGetFromTaskOperationOptions() FileGetFromTaskOperationOptions {
	return FileGetFromTaskOperationOptions{}
}

func (o FileGetFromTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.IfModifiedSince != nil {
		out.Append("If-Modified-Since", fmt.Sprintf("%v", *o.IfModifiedSince))
	}
	if o.IfUnmodifiedSince != nil {
		out.Append("If-Unmodified-Since", fmt.Sprintf("%v", *o.IfUnmodifiedSince))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.OcpRange != nil {
		out.Append("ocp-range", fmt.Sprintf("%v", *o.OcpRange))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o FileGetFromTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FileGetFromTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// FileGetFromTask ...
func (c FilesClient) FileGetFromTask(ctx context.Context, id FileId, options FileGetFromTaskOperationOptions) (result FileGetFromTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
