package keys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckKeysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CheckKeysOperationOptions struct {
	AcceptDatetime *string
	After          *string
	Name           *string
	SyncToken      *string
}

func DefaultCheckKeysOperationOptions() CheckKeysOperationOptions {
	return CheckKeysOperationOptions{}
}

func (o CheckKeysOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o CheckKeysOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CheckKeysOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.After != nil {
		out.Append("After", fmt.Sprintf("%v", *o.After))
	}
	if o.Name != nil {
		out.Append("name", fmt.Sprintf("%v", *o.Name))
	}
	return &out
}

// CheckKeys ...
func (c KeysClient) CheckKeys(ctx context.Context, options CheckKeysOperationOptions) (result CheckKeysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodHead,
		OptionsObject: options,
		Path:          "/keys",
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
