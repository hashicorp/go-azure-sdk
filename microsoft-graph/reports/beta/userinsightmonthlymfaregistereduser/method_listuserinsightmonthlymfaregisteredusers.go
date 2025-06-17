package userinsightmonthlymfaregistereduser

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

type ListUserInsightMonthlyMfaRegisteredUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MfaUserCountMetric
}

type ListUserInsightMonthlyMfaRegisteredUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MfaUserCountMetric
}

type ListUserInsightMonthlyMfaRegisteredUsersOperationOptions struct {
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

func DefaultListUserInsightMonthlyMfaRegisteredUsersOperationOptions() ListUserInsightMonthlyMfaRegisteredUsersOperationOptions {
	return ListUserInsightMonthlyMfaRegisteredUsersOperationOptions{}
}

func (o ListUserInsightMonthlyMfaRegisteredUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlyMfaRegisteredUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightMonthlyMfaRegisteredUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlyMfaRegisteredUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyMfaRegisteredUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyMfaRegisteredUsers - Get mfaRegisteredUsers from reports
func (c UserInsightMonthlyMfaRegisteredUserClient) ListUserInsightMonthlyMfaRegisteredUsers(ctx context.Context, options ListUserInsightMonthlyMfaRegisteredUsersOperationOptions) (result ListUserInsightMonthlyMfaRegisteredUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlyMfaRegisteredUsersCustomPager{},
		Path:          "/reports/userInsights/monthly/mfaRegisteredUsers",
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
		Values *[]beta.MfaUserCountMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyMfaRegisteredUsersComplete retrieves all the results into a single object
func (c UserInsightMonthlyMfaRegisteredUserClient) ListUserInsightMonthlyMfaRegisteredUsersComplete(ctx context.Context, options ListUserInsightMonthlyMfaRegisteredUsersOperationOptions) (ListUserInsightMonthlyMfaRegisteredUsersCompleteResult, error) {
	return c.ListUserInsightMonthlyMfaRegisteredUsersCompleteMatchingPredicate(ctx, options, MfaUserCountMetricOperationPredicate{})
}

// ListUserInsightMonthlyMfaRegisteredUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyMfaRegisteredUserClient) ListUserInsightMonthlyMfaRegisteredUsersCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlyMfaRegisteredUsersOperationOptions, predicate MfaUserCountMetricOperationPredicate) (result ListUserInsightMonthlyMfaRegisteredUsersCompleteResult, err error) {
	items := make([]beta.MfaUserCountMetric, 0)

	resp, err := c.ListUserInsightMonthlyMfaRegisteredUsers(ctx, options)
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

	result = ListUserInsightMonthlyMfaRegisteredUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
