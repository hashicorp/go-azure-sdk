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

type ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsWorkFromAnywhereDevice
}

type ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsWorkFromAnywhereDevice
}

type ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevices ...
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevices(ctx context.Context, id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) (result ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCustomPager{},
		Path:       fmt.Sprintf("%s/metricDevices", id.ID()),
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

// ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesComplete(ctx context.Context, id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId) (ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteMatchingPredicate(ctx, id, UserExperienceAnalyticsWorkFromAnywhereDeviceOperationPredicate{})
}

// ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient) ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricId, predicate UserExperienceAnalyticsWorkFromAnywhereDeviceOperationPredicate) (result ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsWorkFromAnywhereDevice, 0)

	resp, err := c.ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevices(ctx, id)
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

	result = ListUserExperienceAnalyticsWorkFromAnywhereMetricMetricDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
