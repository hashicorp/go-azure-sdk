package datatypes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataType
}

type ListByDataProductCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataType
}

// ListByDataProduct ...
func (c DataTypesClient) ListByDataProduct(ctx context.Context, id DataProductId) (result ListByDataProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/dataTypes", id.ID()),
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
		Values *[]DataType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDataProductComplete retrieves all the results into a single object
func (c DataTypesClient) ListByDataProductComplete(ctx context.Context, id DataProductId) (ListByDataProductCompleteResult, error) {
	return c.ListByDataProductCompleteMatchingPredicate(ctx, id, DataTypeOperationPredicate{})
}

// ListByDataProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataTypesClient) ListByDataProductCompleteMatchingPredicate(ctx context.Context, id DataProductId, predicate DataTypeOperationPredicate) (result ListByDataProductCompleteResult, err error) {
	items := make([]DataType, 0)

	resp, err := c.ListByDataProduct(ctx, id)
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

	result = ListByDataProductCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
