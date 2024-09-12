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

type ListIntentDeviceSettingStateSummariesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListIntentDeviceSettingStateSummariesOperationOptions() ListIntentDeviceSettingStateSummariesOperationOptions {
	return ListIntentDeviceSettingStateSummariesOperationOptions{}
}

func (o ListIntentDeviceSettingStateSummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListIntentDeviceSettingStateSummariesOperationOptions) ToOData() *odata.Query {
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

func (o ListIntentDeviceSettingStateSummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListIntentDeviceSettingStateSummaries - Get deviceSettingStateSummaries from deviceManagement. Collection of settings
// and their states and counts of devices that belong to corresponding state for all settings within the intent
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummaries(ctx context.Context, id beta.DeviceManagementIntentId, options ListIntentDeviceSettingStateSummariesOperationOptions) (result ListIntentDeviceSettingStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListIntentDeviceSettingStateSummariesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceSettingStateSummaries", id.ID()),
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
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummariesComplete(ctx context.Context, id beta.DeviceManagementIntentId, options ListIntentDeviceSettingStateSummariesOperationOptions) (ListIntentDeviceSettingStateSummariesCompleteResult, error) {
	return c.ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate(ctx, id, options, DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate{})
}

// ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentDeviceSettingStateSummaryClient) ListIntentDeviceSettingStateSummariesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementIntentId, options ListIntentDeviceSettingStateSummariesOperationOptions, predicate DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate) (result ListIntentDeviceSettingStateSummariesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntentDeviceSettingStateSummary, 0)

	resp, err := c.ListIntentDeviceSettingStateSummaries(ctx, id, options)
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
