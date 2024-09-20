package userexperienceanalyticsbatteryhealthdeviceperformance

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

type ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthDevicePerformance
}

type ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthDevicePerformance
}

type ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions() ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions {
	return ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions{}
}

func (o ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthDevicePerformances - Get
// userExperienceAnalyticsBatteryHealthDevicePerformance from deviceManagement. User Experience Analytics Battery Health
// Device Performance
func (c UserExperienceAnalyticsBatteryHealthDevicePerformanceClient) ListUserExperienceAnalyticsBatteryHealthDevicePerformances(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions) (result ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthDevicePerformance",
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthDevicePerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthDevicePerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthDevicePerformanceClient) ListUserExperienceAnalyticsBatteryHealthDevicePerformancesComplete(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions) (ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsBatteryHealthDevicePerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthDevicePerformanceClient) ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDevicePerformancesOperationOptions, predicate UserExperienceAnalyticsBatteryHealthDevicePerformanceOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthDevicePerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthDevicePerformances(ctx, options)
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

	result = ListUserExperienceAnalyticsBatteryHealthDevicePerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
