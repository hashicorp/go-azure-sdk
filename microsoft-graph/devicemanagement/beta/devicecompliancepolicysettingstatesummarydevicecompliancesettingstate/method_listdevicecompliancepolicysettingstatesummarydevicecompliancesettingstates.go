package devicecompliancepolicysettingstatesummarydevicecompliancesettingstate

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

type ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceComplianceSettingState
}

type ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceComplianceSettingState
}

type ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStates ...
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStates(ctx context.Context, id DeviceManagementDeviceCompliancePolicySettingStateSummaryId) (result ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCustomPager{},
		Path:       fmt.Sprintf("%s/deviceComplianceSettingStates", id.ID()),
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
		Values *[]beta.DeviceComplianceSettingState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesComplete(ctx context.Context, id DeviceManagementDeviceCompliancePolicySettingStateSummaryId) (ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteResult, error) {
	return c.ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteMatchingPredicate(ctx, id, DeviceComplianceSettingStateOperationPredicate{})
}

// ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDeviceCompliancePolicySettingStateSummaryId, predicate DeviceComplianceSettingStateOperationPredicate) (result ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteResult, err error) {
	items := make([]beta.DeviceComplianceSettingState, 0)

	resp, err := c.ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStates(ctx, id)
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

	result = ListDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
