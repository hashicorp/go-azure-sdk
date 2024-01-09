package syncagents

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLinkedDatabasesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncAgentLinkedDatabase
}

type ListLinkedDatabasesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SyncAgentLinkedDatabase
}

// ListLinkedDatabases ...
func (c SyncAgentsClient) ListLinkedDatabases(ctx context.Context, id SyncAgentId) (result ListLinkedDatabasesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/linkedDatabases", id.ID()),
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
		Values *[]SyncAgentLinkedDatabase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLinkedDatabasesComplete retrieves all the results into a single object
func (c SyncAgentsClient) ListLinkedDatabasesComplete(ctx context.Context, id SyncAgentId) (ListLinkedDatabasesCompleteResult, error) {
	return c.ListLinkedDatabasesCompleteMatchingPredicate(ctx, id, SyncAgentLinkedDatabaseOperationPredicate{})
}

// ListLinkedDatabasesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncAgentsClient) ListLinkedDatabasesCompleteMatchingPredicate(ctx context.Context, id SyncAgentId, predicate SyncAgentLinkedDatabaseOperationPredicate) (result ListLinkedDatabasesCompleteResult, err error) {
	items := make([]SyncAgentLinkedDatabase, 0)

	resp, err := c.ListLinkedDatabases(ctx, id)
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

	result = ListLinkedDatabasesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
