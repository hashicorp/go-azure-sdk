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

type ListLogsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SyncGroupLogProperties
}

type ListLogsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SyncGroupLogProperties
}

type ListLogsOperationOptions struct {
	ContinuationToken *string
	EndTime           *string
	StartTime         *string
	Type              *SyncGroupsType
}

func DefaultListLogsOperationOptions() ListLogsOperationOptions {
	return ListLogsOperationOptions{}
}

func (o ListLogsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLogsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListLogsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ContinuationToken != nil {
		out.Append("continuationToken", fmt.Sprintf("%v", *o.ContinuationToken))
	}
	if o.EndTime != nil {
		out.Append("endTime", fmt.Sprintf("%v", *o.EndTime))
	}
	if o.StartTime != nil {
		out.Append("startTime", fmt.Sprintf("%v", *o.StartTime))
	}
	if o.Type != nil {
		out.Append("type", fmt.Sprintf("%v", *o.Type))
	}
	return &out
}

type ListLogsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListLogsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLogs ...
func (c SyncGroupsClient) ListLogs(ctx context.Context, id SyncGroupId, options ListLogsOperationOptions) (result ListLogsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Pager:         &ListLogsCustomPager{},
		Path:          fmt.Sprintf("%s/logs", id.ID()),
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
		Values *[]SyncGroupLogProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLogsComplete retrieves all the results into a single object
func (c SyncGroupsClient) ListLogsComplete(ctx context.Context, id SyncGroupId, options ListLogsOperationOptions) (ListLogsCompleteResult, error) {
	return c.ListLogsCompleteMatchingPredicate(ctx, id, options, SyncGroupLogPropertiesOperationPredicate{})
}

// ListLogsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SyncGroupsClient) ListLogsCompleteMatchingPredicate(ctx context.Context, id SyncGroupId, options ListLogsOperationOptions, predicate SyncGroupLogPropertiesOperationPredicate) (result ListLogsCompleteResult, err error) {
	items := make([]SyncGroupLogProperties, 0)

	resp, err := c.ListLogs(ctx, id, options)
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

	result = ListLogsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
