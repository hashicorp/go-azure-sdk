package devicecompliancepolicysettingstatesummary

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

type ListDeviceCompliancePolicySettingStateSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceCompliancePolicySettingStateSummary
}

type ListDeviceCompliancePolicySettingStateSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceCompliancePolicySettingStateSummary
}

type ListDeviceCompliancePolicySettingStateSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicySettingStateSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicySettingStateSummaries ...
func (c DeviceCompliancePolicySettingStateSummaryClient) ListDeviceCompliancePolicySettingStateSummaries(ctx context.Context) (result ListDeviceCompliancePolicySettingStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCompliancePolicySettingStateSummariesCustomPager{},
		Path:       "/deviceManagement/deviceCompliancePolicySettingStateSummaries",
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
		Values *[]stable.DeviceCompliancePolicySettingStateSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicySettingStateSummariesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicySettingStateSummaryClient) ListDeviceCompliancePolicySettingStateSummariesComplete(ctx context.Context) (ListDeviceCompliancePolicySettingStateSummariesCompleteResult, error) {
	return c.ListDeviceCompliancePolicySettingStateSummariesCompleteMatchingPredicate(ctx, DeviceCompliancePolicySettingStateSummaryOperationPredicate{})
}

// ListDeviceCompliancePolicySettingStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicySettingStateSummaryClient) ListDeviceCompliancePolicySettingStateSummariesCompleteMatchingPredicate(ctx context.Context, predicate DeviceCompliancePolicySettingStateSummaryOperationPredicate) (result ListDeviceCompliancePolicySettingStateSummariesCompleteResult, err error) {
	items := make([]stable.DeviceCompliancePolicySettingStateSummary, 0)

	resp, err := c.ListDeviceCompliancePolicySettingStateSummaries(ctx)
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

	result = ListDeviceCompliancePolicySettingStateSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
