package pipelineruns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunCancelPipelineRunOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type PipelineRunCancelPipelineRunOperationOptions struct {
	IsRecursive *bool
}

func DefaultPipelineRunCancelPipelineRunOperationOptions() PipelineRunCancelPipelineRunOperationOptions {
	return PipelineRunCancelPipelineRunOperationOptions{}
}

func (o PipelineRunCancelPipelineRunOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PipelineRunCancelPipelineRunOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PipelineRunCancelPipelineRunOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IsRecursive != nil {
		out.Append("isRecursive", fmt.Sprintf("%v", *o.IsRecursive))
	}
	return &out
}

// PipelineRunCancelPipelineRun ...
func (c PipelineRunsClient) PipelineRunCancelPipelineRun(ctx context.Context, id PipelineRunId, options PipelineRunCancelPipelineRunOperationOptions) (result PipelineRunCancelPipelineRunOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/cancel", id.Path()),
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
