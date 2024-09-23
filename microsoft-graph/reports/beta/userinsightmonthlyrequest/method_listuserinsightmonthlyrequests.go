package userinsightmonthlyrequest

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

type ListUserInsightMonthlyRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserRequestsMetric
}

type ListUserInsightMonthlyRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserRequestsMetric
}

type ListUserInsightMonthlyRequestsOperationOptions struct {
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

func DefaultListUserInsightMonthlyRequestsOperationOptions() ListUserInsightMonthlyRequestsOperationOptions {
	return ListUserInsightMonthlyRequestsOperationOptions{}
}

func (o ListUserInsightMonthlyRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlyRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightMonthlyRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlyRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyRequests - List monthly requests. Get a list of monthly user requests on apps registered in
// your tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequests(ctx context.Context, options ListUserInsightMonthlyRequestsOperationOptions) (result ListUserInsightMonthlyRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlyRequestsCustomPager{},
		Path:          "/reports/userInsights/monthly/requests",
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
		Values *[]beta.UserRequestsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyRequestsComplete retrieves all the results into a single object
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequestsComplete(ctx context.Context, options ListUserInsightMonthlyRequestsOperationOptions) (ListUserInsightMonthlyRequestsCompleteResult, error) {
	return c.ListUserInsightMonthlyRequestsCompleteMatchingPredicate(ctx, options, UserRequestsMetricOperationPredicate{})
}

// ListUserInsightMonthlyRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyRequestClient) ListUserInsightMonthlyRequestsCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlyRequestsOperationOptions, predicate UserRequestsMetricOperationPredicate) (result ListUserInsightMonthlyRequestsCompleteResult, err error) {
	items := make([]beta.UserRequestsMetric, 0)

	resp, err := c.ListUserInsightMonthlyRequests(ctx, options)
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

	result = ListUserInsightMonthlyRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
