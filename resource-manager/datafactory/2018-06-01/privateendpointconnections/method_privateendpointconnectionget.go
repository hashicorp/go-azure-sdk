package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PrivateEndpointConnectionResource
}

type PrivateEndpointConnectionGetOperationOptions struct {
	IfNoneMatch *string
}

func DefaultPrivateEndpointConnectionGetOperationOptions() PrivateEndpointConnectionGetOperationOptions {
	return PrivateEndpointConnectionGetOperationOptions{}
}

func (o PrivateEndpointConnectionGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o PrivateEndpointConnectionGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o PrivateEndpointConnectionGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PrivateEndpointConnectionGet ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionGet(ctx context.Context, id PrivateEndpointConnectionId, options PrivateEndpointConnectionGetOperationOptions) (result PrivateEndpointConnectionGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          id.ID(),
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

	var model PrivateEndpointConnectionResource
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
