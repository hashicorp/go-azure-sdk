package databasemigrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlDbGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DatabaseMigrationSqlDb
}

type SqlDbGetOperationOptions struct {
	Expand               *string
	MigrationOperationId *string
}

func DefaultSqlDbGetOperationOptions() SqlDbGetOperationOptions {
	return SqlDbGetOperationOptions{}
}

func (o SqlDbGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlDbGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SqlDbGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.MigrationOperationId != nil {
		out.Append("migrationOperationId", fmt.Sprintf("%v", *o.MigrationOperationId))
	}
	return &out
}

// SqlDbGet ...
func (c DatabaseMigrationsClient) SqlDbGet(ctx context.Context, id DatabaseMigrationId, options SqlDbGetOperationOptions) (result SqlDbGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model DatabaseMigrationSqlDb
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
