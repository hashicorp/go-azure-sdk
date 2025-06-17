package manager

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetManagerRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type GetManagerRefOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Metadata         *odata.Metadata
	RetryFunc        client.RequestRetryFunc
}

func DefaultGetManagerRefOperationOptions() GetManagerRefOperationOptions {
	return GetManagerRefOperationOptions{}
}

func (o GetManagerRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetManagerRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetManagerRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetManagerRef - Get ref of manager from me. The user or contact that is this user's manager. Read-only. Supports
// $expand.
func (c ManagerClient) GetManagerRef(ctx context.Context, options GetManagerRefOperationOptions) (result GetManagerRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/manager/$ref",
		RetryFunc:     options.RetryFunc,
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
