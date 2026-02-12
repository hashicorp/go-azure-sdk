package pools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolEvaluateAutoScaleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AutoScaleRun
}

type PoolEvaluateAutoScaleOperationOptions struct {
	ClientRequestId       *string
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultPoolEvaluateAutoScaleOperationOptions() PoolEvaluateAutoScaleOperationOptions {
	return PoolEvaluateAutoScaleOperationOptions{}
}

func (o PoolEvaluateAutoScaleOperationOptions) ToHeaders() *client.Headers {
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

func (o PoolEvaluateAutoScaleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PoolEvaluateAutoScaleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

// PoolEvaluateAutoScale ...
func (c PoolsClient) PoolEvaluateAutoScale(ctx context.Context, id PoolId, input PoolEvaluateAutoScaleParameter, options PoolEvaluateAutoScaleOperationOptions) (result PoolEvaluateAutoScaleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; odata=minimalmetadata; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/evaluateautoscale", id.Path()),
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

	var model AutoScaleRun
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
