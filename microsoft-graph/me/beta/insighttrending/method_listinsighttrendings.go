package insighttrending

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

type ListInsightTrendingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Trending
}

type ListInsightTrendingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Trending
}

type ListInsightTrendingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListInsightTrendingsOperationOptions() ListInsightTrendingsOperationOptions {
	return ListInsightTrendingsOperationOptions{}
}

func (o ListInsightTrendingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInsightTrendingsOperationOptions) ToOData() *odata.Query {
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

func (o ListInsightTrendingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInsightTrendingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInsightTrendingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInsightTrendings - List trending. Calculated insight that includes a list of documents trending around the user.
func (c InsightTrendingClient) ListInsightTrendings(ctx context.Context, options ListInsightTrendingsOperationOptions) (result ListInsightTrendingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInsightTrendingsCustomPager{},
		Path:          "/me/insights/trending",
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
		Values *[]beta.Trending `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInsightTrendingsComplete retrieves all the results into a single object
func (c InsightTrendingClient) ListInsightTrendingsComplete(ctx context.Context, options ListInsightTrendingsOperationOptions) (ListInsightTrendingsCompleteResult, error) {
	return c.ListInsightTrendingsCompleteMatchingPredicate(ctx, options, TrendingOperationPredicate{})
}

// ListInsightTrendingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InsightTrendingClient) ListInsightTrendingsCompleteMatchingPredicate(ctx context.Context, options ListInsightTrendingsOperationOptions, predicate TrendingOperationPredicate) (result ListInsightTrendingsCompleteResult, err error) {
	items := make([]beta.Trending, 0)

	resp, err := c.ListInsightTrendings(ctx, options)
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

	result = ListInsightTrendingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
