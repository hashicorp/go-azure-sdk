package deviceconfigurationdevicesettingstatesummary

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

type ListDeviceConfigurationDeviceSettingStateSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SettingStateDeviceSummary
}

type ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SettingStateDeviceSummary
}

type ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDeviceConfigurationDeviceSettingStateSummariesOperationOptions() ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions {
	return ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions{}
}

func (o ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListDeviceConfigurationDeviceSettingStateSummaries - List settingStateDeviceSummaries. List properties and
// relationships of the settingStateDeviceSummary objects.
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummaries(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, options ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions) (result ListDeviceConfigurationDeviceSettingStateSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceConfigurationDeviceSettingStateSummariesCustomPager{},
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
		Values *[]stable.SettingStateDeviceSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationDeviceSettingStateSummariesComplete retrieves all the results into a single object
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummariesComplete(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, options ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions) (ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult, error) {
	return c.ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate(ctx, id, options, SettingStateDeviceSummaryOperationPredicate{})
}

// ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationDeviceSettingStateSummaryClient) ListDeviceConfigurationDeviceSettingStateSummariesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, options ListDeviceConfigurationDeviceSettingStateSummariesOperationOptions, predicate SettingStateDeviceSummaryOperationPredicate) (result ListDeviceConfigurationDeviceSettingStateSummariesCompleteResult, err error) {
	items := make([]stable.SettingStateDeviceSummary, 0)

	resp, err := c.ListDeviceConfigurationDeviceSettingStateSummaries(ctx, id, options)
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
