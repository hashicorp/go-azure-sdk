package manageddatabasetables

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
	Model        *[]Resource
}

type ListBySchemaCompleteResult struct {
	Items []Resource
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
func (c ManagedDatabaseTablesClient) ListBySchema(ctx context.Context, id DatabaseSchemaId, options ListBySchemaOperationOptions) (result ListBySchemaOperationResponse, err error) {
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
		Values *[]Resource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySchemaComplete retrieves all the results into a single object
func (c ManagedDatabaseTablesClient) ListBySchemaComplete(ctx context.Context, id DatabaseSchemaId, options ListBySchemaOperationOptions) (ListBySchemaCompleteResult, error) {
	return c.ListBySchemaCompleteMatchingPredicate(ctx, id, options, ResourceOperationPredicate{})
}

// ListBySchemaCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDatabaseTablesClient) ListBySchemaCompleteMatchingPredicate(ctx context.Context, id DatabaseSchemaId, options ListBySchemaOperationOptions, predicate ResourceOperationPredicate) (result ListBySchemaCompleteResult, err error) {
	items := make([]Resource, 0)

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
