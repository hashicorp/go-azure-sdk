package entities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *GetQueriesResponse
}

type QueriesOperationOptions struct {
	Kind *EntityItemQueryKind
}

func DefaultQueriesOperationOptions() QueriesOperationOptions {
	return QueriesOperationOptions{}
}

func (o QueriesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o QueriesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o QueriesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Kind != nil {
		out.Append("kind", fmt.Sprintf("%v", *o.Kind))
	}
	return &out
}

// Queries ...
func (c EntitiesClient) Queries(ctx context.Context, id EntityId, options QueriesOperationOptions) (result QueriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/queries", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
