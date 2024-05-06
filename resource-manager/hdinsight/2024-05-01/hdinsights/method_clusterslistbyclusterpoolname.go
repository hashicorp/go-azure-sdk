package hdinsights

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersListByClusterPoolNameOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Cluster
}

type ClustersListByClusterPoolNameCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Cluster
}

// ClustersListByClusterPoolName ...
func (c HdinsightsClient) ClustersListByClusterPoolName(ctx context.Context, id ClusterPoolId) (result ClustersListByClusterPoolNameOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/clusters", id.ID()),
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
		Values *[]Cluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ClustersListByClusterPoolNameComplete retrieves all the results into a single object
func (c HdinsightsClient) ClustersListByClusterPoolNameComplete(ctx context.Context, id ClusterPoolId) (ClustersListByClusterPoolNameCompleteResult, error) {
	return c.ClustersListByClusterPoolNameCompleteMatchingPredicate(ctx, id, ClusterOperationPredicate{})
}

// ClustersListByClusterPoolNameCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HdinsightsClient) ClustersListByClusterPoolNameCompleteMatchingPredicate(ctx context.Context, id ClusterPoolId, predicate ClusterOperationPredicate) (result ClustersListByClusterPoolNameCompleteResult, err error) {
	items := make([]Cluster, 0)

	resp, err := c.ClustersListByClusterPoolName(ctx, id)
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

	result = ClustersListByClusterPoolNameCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
