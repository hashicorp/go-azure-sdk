package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSupportedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PrivateLinkResourcesResult
}

type ListSupportedOperationOptions struct {
	XMsClientRequestId *string
}

func DefaultListSupportedOperationOptions() ListSupportedOperationOptions {
	return ListSupportedOperationOptions{}
}

func (o ListSupportedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsClientRequestId != nil {
		out.Append("x-ms-client-request-id", fmt.Sprintf("%v", *o.XMsClientRequestId))
	}
	return &out
}

func (o ListSupportedOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListSupportedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ListSupported ...
func (c PrivateLinkResourcesClient) ListSupported(ctx context.Context, id SearchServiceId, options ListSupportedOperationOptions) (result ListSupportedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          fmt.Sprintf("%s/privateLinkResources", id.ID()),
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

	var model PrivateLinkResourcesResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
