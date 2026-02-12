package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineCreatePipelineRunOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateRunResponse
}

type PipelineCreatePipelineRunOperationOptions struct {
	IsRecovery             *bool
	ReferencePipelineRunId *string
	StartActivityName      *string
}

func DefaultPipelineCreatePipelineRunOperationOptions() PipelineCreatePipelineRunOperationOptions {
	return PipelineCreatePipelineRunOperationOptions{}
}

func (o PipelineCreatePipelineRunOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PipelineCreatePipelineRunOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PipelineCreatePipelineRunOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IsRecovery != nil {
		out.Append("isRecovery", fmt.Sprintf("%v", *o.IsRecovery))
	}
	if o.ReferencePipelineRunId != nil {
		out.Append("referencePipelineRunId", fmt.Sprintf("%v", *o.ReferencePipelineRunId))
	}
	if o.StartActivityName != nil {
		out.Append("startActivityName", fmt.Sprintf("%v", *o.StartActivityName))
	}
	return &out
}

// PipelineCreatePipelineRun ...
func (c PipelinesClient) PipelineCreatePipelineRun(ctx context.Context, id PipelineId, input map[string]interface{}, options PipelineCreatePipelineRunOperationOptions) (result PipelineCreatePipelineRunOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createRun", id.Path()),
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

	var model CreateRunResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
