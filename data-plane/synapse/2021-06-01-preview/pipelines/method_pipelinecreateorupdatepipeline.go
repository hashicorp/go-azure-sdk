package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineCreateOrUpdatePipelineOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PipelineResource
}

type PipelineCreateOrUpdatePipelineOperationOptions struct {
	IfMatch *string
}

func DefaultPipelineCreateOrUpdatePipelineOperationOptions() PipelineCreateOrUpdatePipelineOperationOptions {
	return PipelineCreateOrUpdatePipelineOperationOptions{}
}

func (o PipelineCreateOrUpdatePipelineOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o PipelineCreateOrUpdatePipelineOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PipelineCreateOrUpdatePipelineOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PipelineCreateOrUpdatePipeline ...
func (c PipelinesClient) PipelineCreateOrUpdatePipeline(ctx context.Context, id PipelineId, input PipelineResource, options PipelineCreateOrUpdatePipelineOperationOptions) (result PipelineCreateOrUpdatePipelineOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.Path(),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// PipelineCreateOrUpdatePipelineThenPoll performs PipelineCreateOrUpdatePipeline then polls until it's completed
func (c PipelinesClient) PipelineCreateOrUpdatePipelineThenPoll(ctx context.Context, id PipelineId, input PipelineResource, options PipelineCreateOrUpdatePipelineOperationOptions) error {
	result, err := c.PipelineCreateOrUpdatePipeline(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing PipelineCreateOrUpdatePipeline: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PipelineCreateOrUpdatePipeline: %+v", err)
	}

	return nil
}
