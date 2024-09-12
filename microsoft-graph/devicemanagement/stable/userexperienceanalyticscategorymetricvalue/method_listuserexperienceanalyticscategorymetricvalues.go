package userexperienceanalyticscategorymetricvalue

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

type ListUserExperienceAnalyticsCategoryMetricValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsMetric
}

type ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListUserExperienceAnalyticsCategoryMetricValuesOperationOptions() ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions {
	return ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions{}
}

func (o ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsCategoryMetricValuesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsCategoryMetricValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsCategoryMetricValues - Get metricValues from deviceManagement. The metric values for the
// user experience analytics category. Read-only.
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValues(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsCategoryId, options ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions) (result ListUserExperienceAnalyticsCategoryMetricValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsCategoryMetricValuesCustomPager{},
		Path:          fmt.Sprintf("%s/metricValues", id.ID()),
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
		Values *[]stable.UserExperienceAnalyticsMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsCategoryMetricValuesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValuesComplete(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsCategoryId, options ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions) (ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate(ctx, id, options, UserExperienceAnalyticsMetricOperationPredicate{})
}

// ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsCategoryMetricValueClient) ListUserExperienceAnalyticsCategoryMetricValuesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsCategoryId, options ListUserExperienceAnalyticsCategoryMetricValuesOperationOptions, predicate UserExperienceAnalyticsMetricOperationPredicate) (result ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsMetric, 0)

	resp, err := c.ListUserExperienceAnalyticsCategoryMetricValues(ctx, id, options)
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

	result = ListUserExperienceAnalyticsCategoryMetricValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
