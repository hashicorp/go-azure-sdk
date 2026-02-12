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

type CheckKeyValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CheckKeyValuesOperationOptions struct {
	AcceptDatetime *string
	After          *string
	Key            *string
	Label          *string
	Select         *string
	SyncToken      *string
}

func DefaultCheckKeyValuesOperationOptions() CheckKeyValuesOperationOptions {
	return CheckKeyValuesOperationOptions{}
}

func (o CheckKeyValuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o CheckKeyValuesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o CheckKeyValuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.After != nil {
		out.Append("After", fmt.Sprintf("%v", *o.After))
	}
	if o.Key != nil {
		out.Append("key", fmt.Sprintf("%v", *o.Key))
	}
	if o.Label != nil {
		out.Append("label", fmt.Sprintf("%v", *o.Label))
	}
	if o.Select != nil {
		out.Append("$Select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

// CheckKeyValues ...
func (c KeyValuesClient) CheckKeyValues(ctx context.Context, options CheckKeyValuesOperationOptions) (result CheckKeyValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodHead,
		OptionsObject: options,
		Path:          "/kv",
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
