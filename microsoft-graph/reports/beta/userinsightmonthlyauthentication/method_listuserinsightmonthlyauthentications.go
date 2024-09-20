package userinsightmonthlyauthentication

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

type ListUserInsightMonthlyAuthenticationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthenticationsMetric
}

type ListUserInsightMonthlyAuthenticationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthenticationsMetric
}

type ListUserInsightMonthlyAuthenticationsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListUserInsightMonthlyAuthenticationsOperationOptions() ListUserInsightMonthlyAuthenticationsOperationOptions {
	return ListUserInsightMonthlyAuthenticationsOperationOptions{}
}

func (o ListUserInsightMonthlyAuthenticationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlyAuthenticationsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightMonthlyAuthenticationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlyAuthenticationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyAuthenticationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyAuthentications - List monthly authentications. Get a list of monthly authentications on apps
// registered in your tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightMonthlyAuthenticationClient) ListUserInsightMonthlyAuthentications(ctx context.Context, options ListUserInsightMonthlyAuthenticationsOperationOptions) (result ListUserInsightMonthlyAuthenticationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlyAuthenticationsCustomPager{},
		Path:          "/reports/userInsights/monthly/authentications",
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
		Values *[]beta.AuthenticationsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyAuthenticationsComplete retrieves all the results into a single object
func (c UserInsightMonthlyAuthenticationClient) ListUserInsightMonthlyAuthenticationsComplete(ctx context.Context, options ListUserInsightMonthlyAuthenticationsOperationOptions) (ListUserInsightMonthlyAuthenticationsCompleteResult, error) {
	return c.ListUserInsightMonthlyAuthenticationsCompleteMatchingPredicate(ctx, options, AuthenticationsMetricOperationPredicate{})
}

// ListUserInsightMonthlyAuthenticationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyAuthenticationClient) ListUserInsightMonthlyAuthenticationsCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlyAuthenticationsOperationOptions, predicate AuthenticationsMetricOperationPredicate) (result ListUserInsightMonthlyAuthenticationsCompleteResult, err error) {
	items := make([]beta.AuthenticationsMetric, 0)

	resp, err := c.ListUserInsightMonthlyAuthentications(ctx, options)
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

	result = ListUserInsightMonthlyAuthenticationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
