package monthlyprintusagebyuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMonthlyPrintUsageByUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrintUsageByUser
}

type ListMonthlyPrintUsageByUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrintUsageByUser
}

type ListMonthlyPrintUsageByUsersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMonthlyPrintUsageByUsersOperationOptions() ListMonthlyPrintUsageByUsersOperationOptions {
	return ListMonthlyPrintUsageByUsersOperationOptions{}
}

func (o ListMonthlyPrintUsageByUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMonthlyPrintUsageByUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListMonthlyPrintUsageByUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMonthlyPrintUsageByUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMonthlyPrintUsageByUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMonthlyPrintUsageByUsers - List monthlyPrintUsageByUser. Retrieve a list of monthly print usage summaries,
// grouped by user.
func (c MonthlyPrintUsageByUserClient) ListMonthlyPrintUsageByUsers(ctx context.Context, options ListMonthlyPrintUsageByUsersOperationOptions) (result ListMonthlyPrintUsageByUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMonthlyPrintUsageByUsersCustomPager{},
		Path:          "/reports/monthlyPrintUsageByUser",
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
		Values *[]stable.PrintUsageByUser `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMonthlyPrintUsageByUsersComplete retrieves all the results into a single object
func (c MonthlyPrintUsageByUserClient) ListMonthlyPrintUsageByUsersComplete(ctx context.Context, options ListMonthlyPrintUsageByUsersOperationOptions) (ListMonthlyPrintUsageByUsersCompleteResult, error) {
	return c.ListMonthlyPrintUsageByUsersCompleteMatchingPredicate(ctx, options, PrintUsageByUserOperationPredicate{})
}

// ListMonthlyPrintUsageByUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MonthlyPrintUsageByUserClient) ListMonthlyPrintUsageByUsersCompleteMatchingPredicate(ctx context.Context, options ListMonthlyPrintUsageByUsersOperationOptions, predicate PrintUsageByUserOperationPredicate) (result ListMonthlyPrintUsageByUsersCompleteResult, err error) {
	items := make([]stable.PrintUsageByUser, 0)

	resp, err := c.ListMonthlyPrintUsageByUsers(ctx, options)
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

	result = ListMonthlyPrintUsageByUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
