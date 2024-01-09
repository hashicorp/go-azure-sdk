package elasticpooloperations

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

type ListByElasticPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ElasticPoolOperation
}

type ListByElasticPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ElasticPoolOperation
}

// ListByElasticPool ...
func (c ElasticPoolOperationsClient) ListByElasticPool(ctx context.Context, id commonids.SqlElasticPoolId) (result ListByElasticPoolOperationResponse, err error) {
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
		Values *[]ElasticPoolOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByElasticPoolComplete retrieves all the results into a single object
func (c ElasticPoolOperationsClient) ListByElasticPoolComplete(ctx context.Context, id commonids.SqlElasticPoolId) (ListByElasticPoolCompleteResult, error) {
	return c.ListByElasticPoolCompleteMatchingPredicate(ctx, id, ElasticPoolOperationOperationPredicate{})
}

// ListByElasticPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ElasticPoolOperationsClient) ListByElasticPoolCompleteMatchingPredicate(ctx context.Context, id commonids.SqlElasticPoolId, predicate ElasticPoolOperationOperationPredicate) (result ListByElasticPoolCompleteResult, err error) {
	items := make([]ElasticPoolOperation, 0)

	resp, err := c.ListByElasticPool(ctx, id)
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

	result = ListByElasticPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
