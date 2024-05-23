package clusterrecoverypoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByReplicationProtectionClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterRecoveryPoint
}

type ListByReplicationProtectionClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterRecoveryPoint
}

// ListByReplicationProtectionCluster ...
func (c ClusterRecoveryPointsClient) ListByReplicationProtectionCluster(ctx context.Context, id ReplicationProtectionClusterId) (result ListByReplicationProtectionClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/recoveryPoints", id.ID()),
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
		Values *[]ClusterRecoveryPoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByReplicationProtectionClusterComplete retrieves all the results into a single object
func (c ClusterRecoveryPointsClient) ListByReplicationProtectionClusterComplete(ctx context.Context, id ReplicationProtectionClusterId) (ListByReplicationProtectionClusterCompleteResult, error) {
	return c.ListByReplicationProtectionClusterCompleteMatchingPredicate(ctx, id, ClusterRecoveryPointOperationPredicate{})
}

// ListByReplicationProtectionClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClusterRecoveryPointsClient) ListByReplicationProtectionClusterCompleteMatchingPredicate(ctx context.Context, id ReplicationProtectionClusterId, predicate ClusterRecoveryPointOperationPredicate) (result ListByReplicationProtectionClusterCompleteResult, err error) {
	items := make([]ClusterRecoveryPoint, 0)

	resp, err := c.ListByReplicationProtectionCluster(ctx, id)
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

	result = ListByReplicationProtectionClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
