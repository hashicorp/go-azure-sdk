package userexperienceanalyticsworkfromanywheremetricmetricdevice

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

type ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsWorkFromAnywhereDevice
}

type ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsWorkFromAnywhereDevice
}

type ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions() ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions {
	return ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions{}
}

func (o ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsWorkFromAnywhereMetricDevices - Get metricDevices from deviceManagement. The work from
// anywhere metric devices. Read-only.
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricDevices(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, options ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions) (result ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/metricDevices", id.ID()),
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
		Values *[]stable.UserExperienceAnalyticsWorkFromAnywhereDevice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesComplete(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, options ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions) (ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteMatchingPredicate(ctx, id, options, UserExperienceAnalyticsWorkFromAnywhereDeviceOperationPredicate{})
}

// ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, options ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesOperationOptions, predicate UserExperienceAnalyticsWorkFromAnywhereDeviceOperationPredicate) (result ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsWorkFromAnywhereDevice, 0)

	resp, err := c.ListUserExperienceAnalyticsWorkFromAnywhereMetricDevices(ctx, id, options)
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

	result = ListUserExperienceAnalyticsWorkFromAnywhereMetricDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
