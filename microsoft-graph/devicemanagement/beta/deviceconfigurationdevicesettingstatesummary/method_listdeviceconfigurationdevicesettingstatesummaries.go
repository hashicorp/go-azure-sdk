package deviceconfigurationdevicesettingstatesummary

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

type ListDeviceConfigurationDeviceSettingStateSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SettingStateDeviceSummary
}

type ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SettingStateDeviceSummary
}

type ListDeviceConfigurationDeviceSettingStateSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationDeviceSettingStateSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationDeviceSettingStateSummaries ...
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummaries(ctx context.Context, id DeviceManagementDeviceConfigurationId) (result ListDeviceConfigurationDeviceSettingStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationDeviceSettingStateSummariesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceSettingStateSummaries", id.ID()),
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
		Values *[]beta.SettingStateDeviceSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationDeviceSettingStateSummariesComplete retrieves all the results into a single object
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummariesComplete(ctx context.Context, id DeviceManagementDeviceConfigurationId) (ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult, error) {
	return c.ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate(ctx, id, SettingStateDeviceSummaryOperationPredicate{})
}

// ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceConfigurationId, predicate SettingStateDeviceSummaryOperationPredicate) (result ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult, err error) {
	items := make([]beta.SettingStateDeviceSummary, 0)

	resp, err := c.ListDeviceConfigurationDeviceSettingStateSummaries(ctx, id)
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

	result = ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
