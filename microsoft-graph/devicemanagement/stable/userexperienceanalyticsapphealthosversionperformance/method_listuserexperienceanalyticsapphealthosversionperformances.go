package userexperienceanalyticsapphealthosversionperformance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsAppHealthOSVersionPerformance
}

type ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsAppHealthOSVersionPerformance
}

type ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions() ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions {
	return ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions{}
}

func (o ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthOSVersionPerformances - Get userExperienceAnalyticsAppHealthOSVersionPerformance
// from deviceManagement. User experience analytics appHealth OS version Performance
func (c UserExperienceAnalyticsAppHealthOSVersionPerformanceClient) ListUserExperienceAnalyticsAppHealthOSVersionPerformances(ctx context.Context, options ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions) (result ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthOSVersionPerformance",
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
		Values *[]stable.UserExperienceAnalyticsAppHealthOSVersionPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthOSVersionPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthOSVersionPerformanceClient) ListUserExperienceAnalyticsAppHealthOSVersionPerformancesComplete(ctx context.Context, options ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions) (ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAppHealthOSVersionPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthOSVersionPerformanceClient) ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAppHealthOSVersionPerformancesOperationOptions, predicate UserExperienceAnalyticsAppHealthOSVersionPerformanceOperationPredicate) (result ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsAppHealthOSVersionPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthOSVersionPerformances(ctx, options)
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

	result = ListUserExperienceAnalyticsAppHealthOSVersionPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
