package apimanagementserviceskus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAvailableServiceSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceSkuResult
}

type ListAvailableServiceSkusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceSkuResult
}

// ListAvailableServiceSkus ...
func (c ApiManagementServiceSkusClient) ListAvailableServiceSkus(ctx context.Context, id ServiceId) (result ListAvailableServiceSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/skus", id.ID()),
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
		Values *[]ResourceSkuResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAvailableServiceSkusComplete retrieves all the results into a single object
func (c ApiManagementServiceSkusClient) ListAvailableServiceSkusComplete(ctx context.Context, id ServiceId) (ListAvailableServiceSkusCompleteResult, error) {
	return c.ListAvailableServiceSkusCompleteMatchingPredicate(ctx, id, ResourceSkuResultOperationPredicate{})
}

// ListAvailableServiceSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApiManagementServiceSkusClient) ListAvailableServiceSkusCompleteMatchingPredicate(ctx context.Context, id ServiceId, predicate ResourceSkuResultOperationPredicate) (result ListAvailableServiceSkusCompleteResult, err error) {
	items := make([]ResourceSkuResult, 0)

	resp, err := c.ListAvailableServiceSkus(ctx, id)
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

	result = ListAvailableServiceSkusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
