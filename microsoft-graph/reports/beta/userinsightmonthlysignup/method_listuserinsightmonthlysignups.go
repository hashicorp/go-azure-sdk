package userinsightmonthlysignup

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

type ListUserInsightMonthlySignUpsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserSignUpMetric
}

type ListUserInsightMonthlySignUpsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserSignUpMetric
}

type ListUserInsightMonthlySignUpsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListUserInsightMonthlySignUpsOperationOptions() ListUserInsightMonthlySignUpsOperationOptions {
	return ListUserInsightMonthlySignUpsOperationOptions{}
}

func (o ListUserInsightMonthlySignUpsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlySignUpsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightMonthlySignUpsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlySignUpsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlySignUpsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlySignUps - List monthly signUps. Get a list of monthly user sign-ups on apps registered in your
// tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightMonthlySignUpClient) ListUserInsightMonthlySignUps(ctx context.Context, options ListUserInsightMonthlySignUpsOperationOptions) (result ListUserInsightMonthlySignUpsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlySignUpsCustomPager{},
		Path:          "/reports/userInsights/monthly/signUps",
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
		Values *[]beta.UserSignUpMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlySignUpsComplete retrieves all the results into a single object
func (c UserInsightMonthlySignUpClient) ListUserInsightMonthlySignUpsComplete(ctx context.Context, options ListUserInsightMonthlySignUpsOperationOptions) (ListUserInsightMonthlySignUpsCompleteResult, error) {
	return c.ListUserInsightMonthlySignUpsCompleteMatchingPredicate(ctx, options, UserSignUpMetricOperationPredicate{})
}

// ListUserInsightMonthlySignUpsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlySignUpClient) ListUserInsightMonthlySignUpsCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlySignUpsOperationOptions, predicate UserSignUpMetricOperationPredicate) (result ListUserInsightMonthlySignUpsCompleteResult, err error) {
	items := make([]beta.UserSignUpMetric, 0)

	resp, err := c.ListUserInsightMonthlySignUps(ctx, options)
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

	result = ListUserInsightMonthlySignUpsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
