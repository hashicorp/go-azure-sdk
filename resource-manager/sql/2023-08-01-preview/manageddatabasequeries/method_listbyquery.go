package manageddatabasequeries

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByQueryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]QueryStatistics
}

type ListByQueryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []QueryStatistics
}

type ListByQueryOperationOptions struct {
	EndTime   *string
	Interval  *QueryTimeGrainType
	StartTime *string
}

func DefaultListByQueryOperationOptions() ListByQueryOperationOptions {
	return ListByQueryOperationOptions{}
}

func (o ListByQueryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByQueryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByQueryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndTime != nil {
		out.Append("endTime", fmt.Sprintf("%v", *o.EndTime))
	}
	if o.Interval != nil {
		out.Append("interval", fmt.Sprintf("%v", *o.Interval))
	}
	if o.StartTime != nil {
		out.Append("startTime", fmt.Sprintf("%v", *o.StartTime))
	}
	return &out
}

type ListByQueryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByQueryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByQuery ...
func (c ManagedDatabaseQueriesClient) ListByQuery(ctx context.Context, id QueryId, options ListByQueryOperationOptions) (result ListByQueryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByQueryCustomPager{},
		Path:          fmt.Sprintf("%s/statistics", id.ID()),
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
		Values *[]QueryStatistics `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByQueryComplete retrieves all the results into a single object
func (c ManagedDatabaseQueriesClient) ListByQueryComplete(ctx context.Context, id QueryId, options ListByQueryOperationOptions) (ListByQueryCompleteResult, error) {
	return c.ListByQueryCompleteMatchingPredicate(ctx, id, options, QueryStatisticsOperationPredicate{})
}

// ListByQueryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDatabaseQueriesClient) ListByQueryCompleteMatchingPredicate(ctx context.Context, id QueryId, options ListByQueryOperationOptions, predicate QueryStatisticsOperationPredicate) (result ListByQueryCompleteResult, err error) {
	items := make([]QueryStatistics, 0)

	resp, err := c.ListByQuery(ctx, id, options)
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

	result = ListByQueryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
