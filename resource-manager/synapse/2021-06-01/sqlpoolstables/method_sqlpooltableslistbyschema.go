package sqlpoolstables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolTablesListBySchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type SqlPoolTablesListBySchemaCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type SqlPoolTablesListBySchemaOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolTablesListBySchemaOperationOptions() SqlPoolTablesListBySchemaOperationOptions {
	return SqlPoolTablesListBySchemaOperationOptions{}
}

func (o SqlPoolTablesListBySchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolTablesListBySchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolTablesListBySchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type SqlPoolTablesListBySchemaCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolTablesListBySchemaCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolTablesListBySchema ...
func (c SqlPoolsTablesClient) SqlPoolTablesListBySchema(ctx context.Context, id SchemaId, options SqlPoolTablesListBySchemaOperationOptions) (result SqlPoolTablesListBySchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &SqlPoolTablesListBySchemaCustomPager{},
		Path:          fmt.Sprintf("%s/tables", id.ID()),
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

// SqlPoolTablesListBySchemaComplete retrieves all the results into a single object
func (c SqlPoolsTablesClient) SqlPoolTablesListBySchemaComplete(ctx context.Context, id SchemaId, options SqlPoolTablesListBySchemaOperationOptions) (SqlPoolTablesListBySchemaCompleteResult, error) {
	return c.SqlPoolTablesListBySchemaCompleteMatchingPredicate(ctx, id, options, ResourceOperationPredicate{})
}

// SqlPoolTablesListBySchemaCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsTablesClient) SqlPoolTablesListBySchemaCompleteMatchingPredicate(ctx context.Context, id SchemaId, options SqlPoolTablesListBySchemaOperationOptions, predicate ResourceOperationPredicate) (result SqlPoolTablesListBySchemaCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.SqlPoolTablesListBySchema(ctx, id, options)
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

	result = SqlPoolTablesListBySchemaCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
