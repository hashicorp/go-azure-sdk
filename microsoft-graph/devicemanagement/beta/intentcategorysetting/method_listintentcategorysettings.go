package intentcategorysetting

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListIntentCategorySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListIntentCategorySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListIntentCategorySettingsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListIntentCategorySettingsOperationOptions() ListIntentCategorySettingsOperationOptions {
	return ListIntentCategorySettingsOperationOptions{}
}

func (o ListIntentCategorySettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListIntentCategorySettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListIntentCategorySettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListIntentCategorySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentCategorySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentCategorySettings - Get settings from deviceManagement. The settings this category contains
func (c IntentCategorySettingClient) ListIntentCategorySettings(ctx context.Context, id beta.DeviceManagementIntentIdCategoryId, options ListIntentCategorySettingsOperationOptions) (result ListIntentCategorySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListIntentCategorySettingsCustomPager{},
		Path:          fmt.Sprintf("%s/settings", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DeviceManagementSettingInstance, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDeviceManagementSettingInstanceImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DeviceManagementSettingInstance (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListIntentCategorySettingsComplete retrieves all the results into a single object
func (c IntentCategorySettingClient) ListIntentCategorySettingsComplete(ctx context.Context, id beta.DeviceManagementIntentIdCategoryId, options ListIntentCategorySettingsOperationOptions) (ListIntentCategorySettingsCompleteResult, error) {
	return c.ListIntentCategorySettingsCompleteMatchingPredicate(ctx, id, options, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListIntentCategorySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentCategorySettingClient) ListIntentCategorySettingsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementIntentIdCategoryId, options ListIntentCategorySettingsOperationOptions, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListIntentCategorySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListIntentCategorySettings(ctx, id, options)
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

	result = ListIntentCategorySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
