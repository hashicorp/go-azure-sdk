package userinsightdailyusercount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserInsightDailyUserCountsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserCountMetric
}

type ListUserInsightDailyUserCountsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserCountMetric
}

type ListUserInsightDailyUserCountsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListUserInsightDailyUserCountsOperationOptions() ListUserInsightDailyUserCountsOperationOptions {
	return ListUserInsightDailyUserCountsOperationOptions{}
}

func (o ListUserInsightDailyUserCountsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightDailyUserCountsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListUserInsightDailyUserCountsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightDailyUserCountsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyUserCountsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyUserCounts - List daily userCount. Get a list of daily user count on apps registered in your
// tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCounts(ctx context.Context, options ListUserInsightDailyUserCountsOperationOptions) (result ListUserInsightDailyUserCountsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightDailyUserCountsCustomPager{},
		Path:          "/reports/userInsights/daily/userCount",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.UserCountMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyUserCountsComplete retrieves all the results into a single object
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCountsComplete(ctx context.Context, options ListUserInsightDailyUserCountsOperationOptions) (ListUserInsightDailyUserCountsCompleteResult, error) {
	return c.ListUserInsightDailyUserCountsCompleteMatchingPredicate(ctx, options, UserCountMetricOperationPredicate{})
}

// ListUserInsightDailyUserCountsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyUserCountClient) ListUserInsightDailyUserCountsCompleteMatchingPredicate(ctx context.Context, options ListUserInsightDailyUserCountsOperationOptions, predicate UserCountMetricOperationPredicate) (result ListUserInsightDailyUserCountsCompleteResult, err error) {
	items := make([]beta.UserCountMetric, 0)

	resp, err := c.ListUserInsightDailyUserCounts(ctx, options)
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

	result = ListUserInsightDailyUserCountsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
