package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BmcKeySetsListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BmcKeySet
}

type BmcKeySetsListByClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BmcKeySet
}

// BmcKeySetsListByCluster ...
func (c NetworkcloudsClient) BmcKeySetsListByCluster(ctx context.Context, id ClusterId) (result BmcKeySetsListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/bmcKeySets", id.ID()),
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
		Values *[]BmcKeySet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BmcKeySetsListByClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BmcKeySetsListByClusterComplete(ctx context.Context, id ClusterId) (BmcKeySetsListByClusterCompleteResult, error) {
	return c.BmcKeySetsListByClusterCompleteMatchingPredicate(ctx, id, BmcKeySetOperationPredicate{})
}

// BmcKeySetsListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BmcKeySetsListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, predicate BmcKeySetOperationPredicate) (result BmcKeySetsListByClusterCompleteResult, err error) {
	items := make([]BmcKeySet, 0)

	resp, err := c.BmcKeySetsListByCluster(ctx, id)
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

	result = BmcKeySetsListByClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
