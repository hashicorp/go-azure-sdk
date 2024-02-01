package sqlpoolsoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolOperationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlPoolOperation
}

type SqlPoolOperationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlPoolOperation
}

// SqlPoolOperationsList ...
func (c SqlPoolsOperationsClient) SqlPoolOperationsList(ctx context.Context, id SqlPoolId) (result SqlPoolOperationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]SqlPoolOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolOperationsListComplete retrieves all the results into a single object
func (c SqlPoolsOperationsClient) SqlPoolOperationsListComplete(ctx context.Context, id SqlPoolId) (SqlPoolOperationsListCompleteResult, error) {
	return c.SqlPoolOperationsListCompleteMatchingPredicate(ctx, id, SqlPoolOperationOperationPredicate{})
}

// SqlPoolOperationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsOperationsClient) SqlPoolOperationsListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolOperationOperationPredicate) (result SqlPoolOperationsListCompleteResult, err error) {
	items := make([]SqlPoolOperation, 0)

	resp, err := c.SqlPoolOperationsList(ctx, id)
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

	result = SqlPoolOperationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
