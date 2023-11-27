package syncgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSyncDatabaseIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncDatabaseIdProperties
}

type ListSyncDatabaseIdsCompleteResult struct {
	Items []SyncDatabaseIdProperties
}

// ListSyncDatabaseIds ...
func (c SyncGroupsClient) ListSyncDatabaseIds(ctx context.Context, id LocationId) (result ListSyncDatabaseIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/syncDatabaseIds", id.ID()),
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
		Values *[]SyncDatabaseIdProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSyncDatabaseIdsComplete retrieves all the results into a single object
func (c SyncGroupsClient) ListSyncDatabaseIdsComplete(ctx context.Context, id LocationId) (ListSyncDatabaseIdsCompleteResult, error) {
	return c.ListSyncDatabaseIdsCompleteMatchingPredicate(ctx, id, SyncDatabaseIdPropertiesOperationPredicate{})
}

// ListSyncDatabaseIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncGroupsClient) ListSyncDatabaseIdsCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate SyncDatabaseIdPropertiesOperationPredicate) (result ListSyncDatabaseIdsCompleteResult, err error) {
	items := make([]SyncDatabaseIdProperties, 0)

	resp, err := c.ListSyncDatabaseIds(ctx, id)
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

	result = ListSyncDatabaseIdsCompleteResult{
		Items: items,
	}
	return
}
