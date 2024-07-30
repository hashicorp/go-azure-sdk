package updatedevicemanagementmacossoftwareupdateaccountsummarycategorysummarystatesummaries

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

type UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MacOSSoftwareUpdateStateSummary
}

type UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MacOSSoftwareUpdateStateSummary
}

type UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries ...
func (c UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient) UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) (result UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCustomPager{},
		Path:       fmt.Sprintf("%s/updateStateSummaries", id.ID()),
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
		Values *[]beta.MacOSSoftwareUpdateStateSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesComplete retrieves all the results into a single object
func (c UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient) UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesComplete(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId) (UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteResult, error) {
	return c.UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteMatchingPredicate(ctx, id, MacOSSoftwareUpdateStateSummaryOperationPredicate{})
}

// UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesClient) UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryId, predicate MacOSSoftwareUpdateStateSummaryOperationPredicate) (result UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteResult, err error) {
	items := make([]beta.MacOSSoftwareUpdateStateSummary, 0)

	resp, err := c.UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummaries(ctx, id)
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

	result = UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryCategorySummaryStateSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
