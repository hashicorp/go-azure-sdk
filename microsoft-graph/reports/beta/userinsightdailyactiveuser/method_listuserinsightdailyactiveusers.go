package userinsightdailyactiveuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserInsightDailyActiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersMetric
}

type ListUserInsightDailyActiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersMetric
}

type ListUserInsightDailyActiveUsersOperationOptions struct {
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

func DefaultListUserInsightDailyActiveUsersOperationOptions() ListUserInsightDailyActiveUsersOperationOptions {
	return ListUserInsightDailyActiveUsersOperationOptions{}
}

func (o ListUserInsightDailyActiveUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightDailyActiveUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightDailyActiveUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightDailyActiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyActiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyActiveUsers - List daily activeUsers. Get a list of daily active users on apps registered in your
// tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsers(ctx context.Context, options ListUserInsightDailyActiveUsersOperationOptions) (result ListUserInsightDailyActiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightDailyActiveUsersCustomPager{},
		Path:          "/reports/userInsights/daily/activeUsers",
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
		Values *[]beta.ActiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyActiveUsersComplete retrieves all the results into a single object
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsersComplete(ctx context.Context, options ListUserInsightDailyActiveUsersOperationOptions) (ListUserInsightDailyActiveUsersCompleteResult, error) {
	return c.ListUserInsightDailyActiveUsersCompleteMatchingPredicate(ctx, options, ActiveUsersMetricOperationPredicate{})
}

// ListUserInsightDailyActiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyActiveUserClient) ListUserInsightDailyActiveUsersCompleteMatchingPredicate(ctx context.Context, options ListUserInsightDailyActiveUsersOperationOptions, predicate ActiveUsersMetricOperationPredicate) (result ListUserInsightDailyActiveUsersCompleteResult, err error) {
	items := make([]beta.ActiveUsersMetric, 0)

	resp, err := c.ListUserInsightDailyActiveUsers(ctx, options)
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

	result = ListUserInsightDailyActiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
