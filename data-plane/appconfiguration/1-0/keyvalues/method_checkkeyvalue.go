package keyvalues

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckKeyValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CheckKeyValueOperationOptions struct {
	AcceptDatetime *string
	IfMatch        *string
	IfNoneMatch    *string
	Label          *string
	Select         *string
	SyncToken      *string
}

func DefaultCheckKeyValueOperationOptions() CheckKeyValueOperationOptions {
	return CheckKeyValueOperationOptions{}
}

func (o CheckKeyValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
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

func (o CheckKeyValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CheckKeyValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Label != nil {
		out.Append("label", fmt.Sprintf("%v", *o.Label))
	}
	if o.Select != nil {
		out.Append("$Select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

// CheckKeyValue ...
func (c KeyValuesClient) CheckKeyValue(ctx context.Context, id KvId, options CheckKeyValueOperationOptions) (result CheckKeyValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodHead,
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

	return
}
