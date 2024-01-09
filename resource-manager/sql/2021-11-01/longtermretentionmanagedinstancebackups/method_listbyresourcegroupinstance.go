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

type ListByResourceGroupInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstanceLongTermRetentionBackup
}

type ListByResourceGroupInstanceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedInstanceLongTermRetentionBackup
}

type ListByResourceGroupInstanceOperationOptions struct {
	DatabaseState         *DatabaseState
	OnlyLatestPerDatabase *bool
}

func DefaultListByResourceGroupInstanceOperationOptions() ListByResourceGroupInstanceOperationOptions {
	return ListByResourceGroupInstanceOperationOptions{}
}

func (o ListByResourceGroupInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByResourceGroupInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByResourceGroupInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatabaseState != nil {
		out.Append("databaseState", fmt.Sprintf("%v", *o.DatabaseState))
	}
	if o.OnlyLatestPerDatabase != nil {
		out.Append("onlyLatestPerDatabase", fmt.Sprintf("%v", *o.OnlyLatestPerDatabase))
	}
	return &out
}

// ListByResourceGroupInstance ...
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupInstance(ctx context.Context, id LocationLongTermRetentionManagedInstanceId, options ListByResourceGroupInstanceOperationOptions) (result ListByResourceGroupInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/longTermRetentionManagedInstanceBackups", id.ID()),
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
		Values *[]ManagedInstanceLongTermRetentionBackup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByResourceGroupInstanceComplete retrieves all the results into a single object
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupInstanceComplete(ctx context.Context, id LocationLongTermRetentionManagedInstanceId, options ListByResourceGroupInstanceOperationOptions) (ListByResourceGroupInstanceCompleteResult, error) {
	return c.ListByResourceGroupInstanceCompleteMatchingPredicate(ctx, id, options, ManagedInstanceLongTermRetentionBackupOperationPredicate{})
}

// ListByResourceGroupInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LongTermRetentionManagedInstanceBackupsClient) ListByResourceGroupInstanceCompleteMatchingPredicate(ctx context.Context, id LocationLongTermRetentionManagedInstanceId, options ListByResourceGroupInstanceOperationOptions, predicate ManagedInstanceLongTermRetentionBackupOperationPredicate) (result ListByResourceGroupInstanceCompleteResult, err error) {
	items := make([]ManagedInstanceLongTermRetentionBackup, 0)

	resp, err := c.ListByResourceGroupInstance(ctx, id, options)
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

	result = ListByResourceGroupInstanceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
