package databasetables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseTable
}

type ListBySchemaCompleteResult struct {
	Items []DatabaseTable
}

type ListBySchemaOperationOptions struct {
	Filter *string
}

func DefaultListBySchemaOperationOptions() ListBySchemaOperationOptions {
	return ListBySchemaOperationOptions{}
}

func (o ListBySchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListBySchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListBySchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListBySchema ...
func (c DatabaseTablesClient) ListBySchema(ctx context.Context, id SchemaId, options ListBySchemaOperationOptions) (result ListBySchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/tables", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]DatabaseTable `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySchemaComplete retrieves all the results into a single object
func (c DatabaseTablesClient) ListBySchemaComplete(ctx context.Context, id SchemaId, options ListBySchemaOperationOptions) (ListBySchemaCompleteResult, error) {
	return c.ListBySchemaCompleteMatchingPredicate(ctx, id, options, DatabaseTableOperationPredicate{})
}

// ListBySchemaCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabaseTablesClient) ListBySchemaCompleteMatchingPredicate(ctx context.Context, id SchemaId, options ListBySchemaOperationOptions, predicate DatabaseTableOperationPredicate) (result ListBySchemaCompleteResult, err error) {
	items := make([]DatabaseTable, 0)

	resp, err := c.ListBySchema(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListBySchemaCompleteResult{
		Items: items,
	}
	return
}
