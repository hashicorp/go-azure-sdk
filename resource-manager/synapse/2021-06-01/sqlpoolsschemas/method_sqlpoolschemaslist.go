package sqlpoolsschemas

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSchemasListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type SqlPoolSchemasListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type SqlPoolSchemasListOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolSchemasListOperationOptions() SqlPoolSchemasListOperationOptions {
	return SqlPoolSchemasListOperationOptions{}
}

func (o SqlPoolSchemasListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolSchemasListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolSchemasListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// SqlPoolSchemasList ...
func (c SqlPoolsSchemasClient) SqlPoolSchemasList(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions) (result SqlPoolSchemasListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/schemas", id.ID()),
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

// SqlPoolSchemasListComplete retrieves all the results into a single object
func (c SqlPoolsSchemasClient) SqlPoolSchemasListComplete(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions) (SqlPoolSchemasListCompleteResult, error) {
	return c.SqlPoolSchemasListCompleteMatchingPredicate(ctx, id, options, ResourceOperationPredicate{})
}

// SqlPoolSchemasListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsSchemasClient) SqlPoolSchemasListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions, predicate ResourceOperationPredicate) (result SqlPoolSchemasListCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.SqlPoolSchemasList(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = SqlPoolSchemasListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
