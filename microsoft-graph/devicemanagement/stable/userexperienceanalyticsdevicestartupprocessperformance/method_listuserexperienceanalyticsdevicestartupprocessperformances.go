package userexperienceanalyticsdevicestartupprocessperformance

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

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance
}

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDeviceStartupProcessPerformance
}

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions() ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions {
	return ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions{}
}

func (o ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformances - Get
// userExperienceAnalyticsDeviceStartupProcessPerformance from deviceManagement. User experience analytics device
// Startup Process Performance
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformances(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions) (result ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance",
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
		Values *[]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformancesComplete(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions) (ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsDeviceStartupProcessPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceStartupProcessPerformanceClient) ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessPerformancesOperationOptions, predicate UserExperienceAnalyticsDeviceStartupProcessPerformanceOperationPredicate) (result ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDeviceStartupProcessPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceStartupProcessPerformances(ctx, options)
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

	result = ListUserExperienceAnalyticsDeviceStartupProcessPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
