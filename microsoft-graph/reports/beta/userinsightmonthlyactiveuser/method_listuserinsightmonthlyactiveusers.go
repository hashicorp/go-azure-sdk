package userinsightmonthlyactiveuser

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

type ListUserInsightMonthlyActiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActiveUsersMetric
}

type ListUserInsightMonthlyActiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActiveUsersMetric
}

type ListUserInsightMonthlyActiveUsersOperationOptions struct {
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

func DefaultListUserInsightMonthlyActiveUsersOperationOptions() ListUserInsightMonthlyActiveUsersOperationOptions {
	return ListUserInsightMonthlyActiveUsersOperationOptions{}
}

func (o ListUserInsightMonthlyActiveUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlyActiveUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightMonthlyActiveUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlyActiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyActiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyActiveUsers - List monthly activeUsers. Get a list of monthly active users on apps registered
// in your tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsers(ctx context.Context, options ListUserInsightMonthlyActiveUsersOperationOptions) (result ListUserInsightMonthlyActiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlyActiveUsersCustomPager{},
		Path:          "/reports/userInsights/monthly/activeUsers",
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

// ListUserInsightMonthlyActiveUsersComplete retrieves all the results into a single object
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsersComplete(ctx context.Context, options ListUserInsightMonthlyActiveUsersOperationOptions) (ListUserInsightMonthlyActiveUsersCompleteResult, error) {
	return c.ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate(ctx, options, ActiveUsersMetricOperationPredicate{})
}

// ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyActiveUserClient) ListUserInsightMonthlyActiveUsersCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlyActiveUsersOperationOptions, predicate ActiveUsersMetricOperationPredicate) (result ListUserInsightMonthlyActiveUsersCompleteResult, err error) {
	items := make([]beta.ActiveUsersMetric, 0)

	resp, err := c.ListUserInsightMonthlyActiveUsers(ctx, options)
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

	result = ListUserInsightMonthlyActiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
