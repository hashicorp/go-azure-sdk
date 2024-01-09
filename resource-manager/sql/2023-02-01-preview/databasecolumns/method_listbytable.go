package databasecolumns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTableOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseColumn
}

type ListByTableCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseColumn
}

type ListByTableOperationOptions struct {
	Filter *string
}

func DefaultListByTableOperationOptions() ListByTableOperationOptions {
	return ListByTableOperationOptions{}
}

func (o ListByTableOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByTableOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByTableOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByTable ...
func (c DatabaseColumnsClient) ListByTable(ctx context.Context, id TableId, options ListByTableOperationOptions) (result ListByTableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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
		Values *[]DatabaseColumn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByTableComplete retrieves all the results into a single object
func (c DatabaseColumnsClient) ListByTableComplete(ctx context.Context, id TableId, options ListByTableOperationOptions) (ListByTableCompleteResult, error) {
	return c.ListByTableCompleteMatchingPredicate(ctx, id, options, DatabaseColumnOperationPredicate{})
}

// ListByTableCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabaseColumnsClient) ListByTableCompleteMatchingPredicate(ctx context.Context, id TableId, options ListByTableOperationOptions, predicate DatabaseColumnOperationPredicate) (result ListByTableCompleteResult, err error) {
	items := make([]DatabaseColumn, 0)

	resp, err := c.ListByTable(ctx, id, options)
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

	result = ListByTableCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
