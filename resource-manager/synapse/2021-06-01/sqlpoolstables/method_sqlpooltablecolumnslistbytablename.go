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

type SqlPoolTableColumnsListByTableNameOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlPoolColumn
}

type SqlPoolTableColumnsListByTableNameCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlPoolColumn
}

type SqlPoolTableColumnsListByTableNameOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolTableColumnsListByTableNameOperationOptions() SqlPoolTableColumnsListByTableNameOperationOptions {
	return SqlPoolTableColumnsListByTableNameOperationOptions{}
}

func (o SqlPoolTableColumnsListByTableNameOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SqlPoolTableColumnsListByTableNameOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o SqlPoolTableColumnsListByTableNameOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type SqlPoolTableColumnsListByTableNameCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolTableColumnsListByTableNameCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolTableColumnsListByTableName ...
func (c SqlPoolsTablesClient) SqlPoolTableColumnsListByTableName(ctx context.Context, id TableId, options SqlPoolTableColumnsListByTableNameOperationOptions) (result SqlPoolTableColumnsListByTableNameOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &SqlPoolTableColumnsListByTableNameCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]SqlPoolColumn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolTableColumnsListByTableNameComplete retrieves all the results into a single object
func (c SqlPoolsTablesClient) SqlPoolTableColumnsListByTableNameComplete(ctx context.Context, id TableId, options SqlPoolTableColumnsListByTableNameOperationOptions) (SqlPoolTableColumnsListByTableNameCompleteResult, error) {
	return c.SqlPoolTableColumnsListByTableNameCompleteMatchingPredicate(ctx, id, options, SqlPoolColumnOperationPredicate{})
}

// SqlPoolTableColumnsListByTableNameCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsTablesClient) SqlPoolTableColumnsListByTableNameCompleteMatchingPredicate(ctx context.Context, id TableId, options SqlPoolTableColumnsListByTableNameOperationOptions, predicate SqlPoolColumnOperationPredicate) (result SqlPoolTableColumnsListByTableNameCompleteResult, err error) {
	items := make([]SqlPoolColumn, 0)

	resp, err := c.SqlPoolTableColumnsListByTableName(ctx, id, options)
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

	result = SqlPoolTableColumnsListByTableNameCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
