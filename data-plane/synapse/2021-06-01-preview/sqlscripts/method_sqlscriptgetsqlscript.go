package sqlscripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlScriptGetSqlScriptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SqlScriptResource
}

type SqlScriptGetSqlScriptOperationOptions struct {
	IfNoneMatch *string
}

func DefaultSqlScriptGetSqlScriptOperationOptions() SqlScriptGetSqlScriptOperationOptions {
	return SqlScriptGetSqlScriptOperationOptions{}
}

func (o SqlScriptGetSqlScriptOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o SqlScriptGetSqlScriptOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SqlScriptGetSqlScriptOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SqlScriptGetSqlScript ...
func (c SqlScriptsClient) SqlScriptGetSqlScript(ctx context.Context, id SqlScriptId, options SqlScriptGetSqlScriptOperationOptions) (result SqlScriptGetSqlScriptOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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

	var model SqlScriptResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
