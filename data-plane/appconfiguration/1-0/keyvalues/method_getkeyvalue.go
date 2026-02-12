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

type GetKeyValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *KeyValue
}

type GetKeyValueOperationOptions struct {
	AcceptDatetime *string
	IfMatch        *string
	IfNoneMatch    *string
	Label          *string
	Select         *string
	SyncToken      *string
}

func DefaultGetKeyValueOperationOptions() GetKeyValueOperationOptions {
	return GetKeyValueOperationOptions{}
}

func (o GetKeyValueOperationOptions) ToHeaders() *client.Headers {
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

func (o GetKeyValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetKeyValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Label != nil {
		out.Append("label", fmt.Sprintf("%v", *o.Label))
	}
	if o.Select != nil {
		out.Append("$Select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

// GetKeyValue ...
func (c KeyValuesClient) GetKeyValue(ctx context.Context, id KvId, options GetKeyValueOperationOptions) (result GetKeyValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/vnd.microsoft.appconfig.kv+json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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
