package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlResourcesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RestorableSqlResourcesListResult
}

type RestorableSqlResourcesListOperationOptions struct {
	RestoreLocation       *string
	RestoreTimestampInUtc *string
}

func DefaultRestorableSqlResourcesListOperationOptions() RestorableSqlResourcesListOperationOptions {
	return RestorableSqlResourcesListOperationOptions{}
}

func (o RestorableSqlResourcesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestorableSqlResourcesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o RestorableSqlResourcesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RestoreLocation != nil {
		out.Append("restoreLocation", fmt.Sprintf("%v", *o.RestoreLocation))
	}
	if o.RestoreTimestampInUtc != nil {
		out.Append("restoreTimestampInUtc", fmt.Sprintf("%v", *o.RestoreTimestampInUtc))
	}
	return &out
}

// RestorableSqlResourcesList ...
func (c RestorablesClient) RestorableSqlResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableSqlResourcesListOperationOptions) (result RestorableSqlResourcesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/restorableSqlResources", id.ID()),
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

	var model RestorableSqlResourcesListResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
