package longtermretentionmanagedinstancebackups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByResourceGroupDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstanceLongTermRetentionBackup
}

type ListByResourceGroupDatabaseCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedInstanceLongTermRetentionBackup
}

type ListByResourceGroupDatabaseOperationOptions struct {
	DatabaseState         *DatabaseState
	OnlyLatestPerDatabase *bool
}

func DefaultListByResourceGroupDatabaseOperationOptions() ListByResourceGroupDatabaseOperationOptions {
	return ListByResourceGroupDatabaseOperationOptions{}
}

func (o ListByResourceGroupDatabaseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByResourceGroupDatabaseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByResourceGroupDatabaseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatabaseState != nil {
		out.Append("databaseState", fmt.Sprintf("%v", *o.DatabaseState))
	}
	if o.OnlyLatestPerDatabase != nil {
		out.Append("onlyLatestPerDatabase", fmt.Sprintf("%v", *o.OnlyLatestPerDatabase))
	}
	return &out
}

type ListByResourceGroupDatabaseCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByResourceGroupDatabaseCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByResourceGroupDatabase ...
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupDatabase(ctx context.Context, id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId, options ListByResourceGroupDatabaseOperationOptions) (result ListByResourceGroupDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByResourceGroupDatabaseCustomPager{},
		Path:          fmt.Sprintf("%s/longTermRetentionManagedInstanceBackups", id.ID()),
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
		Values *[]ManagedInstanceLongTermRetentionBackup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByResourceGroupDatabaseComplete retrieves all the results into a single object
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupDatabaseComplete(ctx context.Context, id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId, options ListByResourceGroupDatabaseOperationOptions) (ListByResourceGroupDatabaseCompleteResult, error) {
	return c.ListByResourceGroupDatabaseCompleteMatchingPredicate(ctx, id, options, ManagedInstanceLongTermRetentionBackupOperationPredicate{})
}

// ListByResourceGroupDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupDatabaseCompleteMatchingPredicate(ctx context.Context, id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId, options ListByResourceGroupDatabaseOperationOptions, predicate ManagedInstanceLongTermRetentionBackupOperationPredicate) (result ListByResourceGroupDatabaseCompleteResult, err error) {
	items := make([]ManagedInstanceLongTermRetentionBackup, 0)

	resp, err := c.ListByResourceGroupDatabase(ctx, id, options)
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

	result = ListByResourceGroupDatabaseCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
