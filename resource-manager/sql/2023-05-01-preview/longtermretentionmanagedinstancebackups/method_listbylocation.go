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

type ListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstanceLongTermRetentionBackup
}

type ListByLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedInstanceLongTermRetentionBackup
}

type ListByLocationOperationOptions struct {
	DatabaseState         *DatabaseState
	Filter                *string
	OnlyLatestPerDatabase *bool
	Skip                  *int64
	Top                   *int64
}

func DefaultListByLocationOperationOptions() ListByLocationOperationOptions {
	return ListByLocationOperationOptions{}
}

func (o ListByLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DatabaseState != nil {
		out.Append("databaseState", fmt.Sprintf("%v", *o.DatabaseState))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OnlyLatestPerDatabase != nil {
		out.Append("onlyLatestPerDatabase", fmt.Sprintf("%v", *o.OnlyLatestPerDatabase))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByLocation ...
func (c LongTermRetentionManagedInstanceBackupsClient) ListByLocation(ctx context.Context, id LocationId, options ListByLocationOperationOptions) (result ListByLocationOperationResponse, err error) {
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

// ListByLocationComplete retrieves all the results into a single object
func (c LongTermRetentionManagedInstanceBackupsClient) ListByLocationComplete(ctx context.Context, id LocationId, options ListByLocationOperationOptions) (ListByLocationCompleteResult, error) {
	return c.ListByLocationCompleteMatchingPredicate(ctx, id, options, ManagedInstanceLongTermRetentionBackupOperationPredicate{})
}

// ListByLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LongTermRetentionManagedInstanceBackupsClient) ListByLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListByLocationOperationOptions, predicate ManagedInstanceLongTermRetentionBackupOperationPredicate) (result ListByLocationCompleteResult, err error) {
	items := make([]ManagedInstanceLongTermRetentionBackup, 0)

	resp, err := c.ListByLocation(ctx, id, options)
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

	result = ListByLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
