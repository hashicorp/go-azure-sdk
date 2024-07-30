package intentdevicesettingstatesummary

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

type ListIntentDeviceSettingStateSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntentDeviceSettingStateSummary
}

type ListIntentDeviceSettingStateSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntentDeviceSettingStateSummary
}

type ListIntentDeviceSettingStateSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentDeviceSettingStateSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentDeviceSettingStateSummaries ...
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummaries(ctx context.Context, id DeviceManagementIntentId) (result ListIntentDeviceSettingStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentDeviceSettingStateSummariesCustomPager{},
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
		Values *[]beta.DeviceManagementIntentDeviceSettingStateSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentDeviceSettingStateSummariesComplete retrieves all the results into a single object
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummariesComplete(ctx context.Context, id DeviceManagementIntentId) (ListIntentDeviceSettingStateSummariesCompleteResult, error) {
	return c.ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate(ctx, id, DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate{})
}

// ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentId, predicate DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate) (result ListIntentDeviceSettingStateSummariesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntentDeviceSettingStateSummary, 0)

	resp, err := c.ListIntentDeviceSettingStateSummaries(ctx, id)
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

	result = ListIntentDeviceSettingStateSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
