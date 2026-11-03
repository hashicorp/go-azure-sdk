package databasecolumns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseColumn
}

type ListByDatabaseCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseColumn
}

type ListByDatabaseOperationOptions struct {
	Column  *[]string
	OrderBy *[]string
	Schema  *[]string
	Table   *[]string
}

func DefaultListByDatabaseOperationOptions() ListByDatabaseOperationOptions {
	return ListByDatabaseOperationOptions{}
}

func (o ListByDatabaseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByDatabaseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByDatabaseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Column != nil {
		out.Append("column", fmt.Sprintf("%v", *o.Column))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Schema != nil {
		out.Append("schema", fmt.Sprintf("%v", *o.Schema))
	}
	if o.Table != nil {
		out.Append("table", fmt.Sprintf("%v", *o.Table))
	}
	return &out
}

type ListByDatabaseCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDatabaseCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDatabase ...
func (c DatabaseColumnsClient) ListByDatabase(ctx context.Context, id commonids.SqlDatabaseId, options ListByDatabaseOperationOptions) (result ListByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByDatabaseCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]DatabaseColumn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDatabaseComplete retrieves all the results into a single object
func (c DatabaseColumnsClient) ListByDatabaseComplete(ctx context.Context, id commonids.SqlDatabaseId, options ListByDatabaseOperationOptions) (ListByDatabaseCompleteResult, error) {
	return c.ListByDatabaseCompleteMatchingPredicate(ctx, id, options, DatabaseColumnOperationPredicate{})
}

// ListByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabaseColumnsClient) ListByDatabaseCompleteMatchingPredicate(ctx context.Context, id commonids.SqlDatabaseId, options ListByDatabaseOperationOptions, predicate DatabaseColumnOperationPredicate) (result ListByDatabaseCompleteResult, err error) {
	items := make([]DatabaseColumn, 0)

	resp, err := c.ListByDatabase(ctx, id, options)
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

	result = ListByDatabaseCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
