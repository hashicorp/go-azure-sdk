package userinsightdailyinactiveuser

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

type ListUserInsightDailyInactiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DailyInactiveUsersMetric
}

type ListUserInsightDailyInactiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DailyInactiveUsersMetric
}

type ListUserInsightDailyInactiveUsersOperationOptions struct {
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

func DefaultListUserInsightDailyInactiveUsersOperationOptions() ListUserInsightDailyInactiveUsersOperationOptions {
	return ListUserInsightDailyInactiveUsersOperationOptions{}
}

func (o ListUserInsightDailyInactiveUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightDailyInactiveUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightDailyInactiveUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightDailyInactiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyInactiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyInactiveUsers - Get inactiveUsers from reports
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsers(ctx context.Context, options ListUserInsightDailyInactiveUsersOperationOptions) (result ListUserInsightDailyInactiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightDailyInactiveUsersCustomPager{},
		Path:          "/reports/userInsights/daily/inactiveUsers",
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
		Values *[]beta.DailyInactiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyInactiveUsersComplete retrieves all the results into a single object
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsersComplete(ctx context.Context, options ListUserInsightDailyInactiveUsersOperationOptions) (ListUserInsightDailyInactiveUsersCompleteResult, error) {
	return c.ListUserInsightDailyInactiveUsersCompleteMatchingPredicate(ctx, options, DailyInactiveUsersMetricOperationPredicate{})
}

// ListUserInsightDailyInactiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyInactiveUserClient) ListUserInsightDailyInactiveUsersCompleteMatchingPredicate(ctx context.Context, options ListUserInsightDailyInactiveUsersOperationOptions, predicate DailyInactiveUsersMetricOperationPredicate) (result ListUserInsightDailyInactiveUsersCompleteResult, err error) {
	items := make([]beta.DailyInactiveUsersMetric, 0)

	resp, err := c.ListUserInsightDailyInactiveUsers(ctx, options)
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

	result = ListUserInsightDailyInactiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
