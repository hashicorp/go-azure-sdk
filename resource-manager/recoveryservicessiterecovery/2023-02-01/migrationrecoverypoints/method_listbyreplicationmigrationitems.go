package migrationrecoverypoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByReplicationMigrationItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MigrationRecoveryPoint
}

type ListByReplicationMigrationItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MigrationRecoveryPoint
}

// ListByReplicationMigrationItems ...
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItems(ctx context.Context, id ReplicationMigrationItemId) (result ListByReplicationMigrationItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/migrationRecoveryPoints", id.ID()),
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
		Values *[]MigrationRecoveryPoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByReplicationMigrationItemsComplete retrieves all the results into a single object
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItemsComplete(ctx context.Context, id ReplicationMigrationItemId) (ListByReplicationMigrationItemsCompleteResult, error) {
	return c.ListByReplicationMigrationItemsCompleteMatchingPredicate(ctx, id, MigrationRecoveryPointOperationPredicate{})
}

// ListByReplicationMigrationItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItemsCompleteMatchingPredicate(ctx context.Context, id ReplicationMigrationItemId, predicate MigrationRecoveryPointOperationPredicate) (result ListByReplicationMigrationItemsCompleteResult, err error) {
	items := make([]MigrationRecoveryPoint, 0)

	resp, err := c.ListByReplicationMigrationItems(ctx, id)
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

	result = ListByReplicationMigrationItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
