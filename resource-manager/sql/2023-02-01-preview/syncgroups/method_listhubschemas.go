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

type ListHubSchemasOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncFullSchemaProperties
}

type ListHubSchemasCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SyncFullSchemaProperties
}

// ListHubSchemas ...
func (c SyncGroupsClient) ListHubSchemas(ctx context.Context, id SyncGroupId) (result ListHubSchemasOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/hubSchemas", id.ID()),
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
		Values *[]SyncFullSchemaProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHubSchemasComplete retrieves all the results into a single object
func (c SyncGroupsClient) ListHubSchemasComplete(ctx context.Context, id SyncGroupId) (ListHubSchemasCompleteResult, error) {
	return c.ListHubSchemasCompleteMatchingPredicate(ctx, id, SyncFullSchemaPropertiesOperationPredicate{})
}

// ListHubSchemasCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncGroupsClient) ListHubSchemasCompleteMatchingPredicate(ctx context.Context, id SyncGroupId, predicate SyncFullSchemaPropertiesOperationPredicate) (result ListHubSchemasCompleteResult, err error) {
	items := make([]SyncFullSchemaProperties, 0)

	resp, err := c.ListHubSchemas(ctx, id)
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

	result = ListHubSchemasCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
