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

type CancelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelOperationOptions struct {
	IsRecursive *bool
}

func DefaultCancelOperationOptions() CancelOperationOptions {
	return CancelOperationOptions{}
}

func (o CancelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o CancelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IsRecursive != nil {
		out.Append("isRecursive", fmt.Sprintf("%v", *o.IsRecursive))
	}
	return &out
}

// Cancel ...
func (c PipelineRunsClient) Cancel(ctx context.Context, id PipelineRunId, options CancelOperationOptions) (result CancelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
		OptionsObject: options,
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
