package userexperienceanalyticsapphealthdevicemodelperformance

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

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance
}

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions() ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions {
	return ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions{}
}

func (o ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformances - Get
// userExperienceAnalyticsAppHealthDeviceModelPerformance from deviceManagement. User experience analytics appHealth
// Model Performance
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformances(ctx context.Context, options ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions) (result ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance",
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
		Values *[]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesComplete(ctx context.Context, options ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions) (ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient) ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesOperationOptions, predicate UserExperienceAnalyticsAppHealthDeviceModelPerformanceOperationPredicate) (result ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAppHealthDeviceModelPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthDeviceModelPerformances(ctx, options)
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

	result = ListUserExperienceAnalyticsAppHealthDeviceModelPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
