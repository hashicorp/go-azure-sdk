package replicationmigrationitems

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByReplicationProtectionContainersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MigrationItem
}

type ListByReplicationProtectionContainersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MigrationItem
}

type ListByReplicationProtectionContainersOperationOptions struct {
	Filter    *string
	SkipToken *string
	TakeToken *string
}

func DefaultListByReplicationProtectionContainersOperationOptions() ListByReplicationProtectionContainersOperationOptions {
	return ListByReplicationProtectionContainersOperationOptions{}
}

func (o ListByReplicationProtectionContainersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByReplicationProtectionContainersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByReplicationProtectionContainersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.SkipToken != nil {
		out.Append("skipToken", fmt.Sprintf("%v", *o.SkipToken))
	}
	if o.TakeToken != nil {
		out.Append("takeToken", fmt.Sprintf("%v", *o.TakeToken))
	}
	return &out
}

// ListByReplicationProtectionContainers ...
func (c ReplicationMigrationItemsClient) ListByReplicationProtectionContainers(ctx context.Context, id ReplicationProtectionContainerId, options ListByReplicationProtectionContainersOperationOptions) (result ListByReplicationProtectionContainersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/replicationMigrationItems", id.ID()),
		OptionsObject: options,
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
		Values *[]MigrationItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByReplicationProtectionContainersComplete retrieves all the results into a single object
func (c ReplicationMigrationItemsClient) ListByReplicationProtectionContainersComplete(ctx context.Context, id ReplicationProtectionContainerId, options ListByReplicationProtectionContainersOperationOptions) (ListByReplicationProtectionContainersCompleteResult, error) {
	return c.ListByReplicationProtectionContainersCompleteMatchingPredicate(ctx, id, options, MigrationItemOperationPredicate{})
}

// ListByReplicationProtectionContainersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReplicationMigrationItemsClient) ListByReplicationProtectionContainersCompleteMatchingPredicate(ctx context.Context, id ReplicationProtectionContainerId, options ListByReplicationProtectionContainersOperationOptions, predicate MigrationItemOperationPredicate) (result ListByReplicationProtectionContainersCompleteResult, err error) {
	items := make([]MigrationItem, 0)

	resp, err := c.ListByReplicationProtectionContainers(ctx, id, options)
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

	result = ListByReplicationProtectionContainersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
