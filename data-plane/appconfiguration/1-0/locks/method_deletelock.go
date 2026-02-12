package locks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteLockOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *KeyValue
}

type DeleteLockOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
	Label       *string
	SyncToken   *string
}

func DefaultDeleteLockOperationOptions() DeleteLockOperationOptions {
	return DeleteLockOperationOptions{}
}

func (o DeleteLockOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o DeleteLockOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteLockOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Label != nil {
		out.Append("label", fmt.Sprintf("%v", *o.Label))
	}
	return &out
}

// DeleteLock ...
func (c LocksClient) DeleteLock(ctx context.Context, id LockId, options DeleteLockOperationOptions) (result DeleteLockOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
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

	var model KeyValue
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
