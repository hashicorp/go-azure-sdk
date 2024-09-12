package templatecategoryrecommendedsetting

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

type ListTemplateCategoryRecommendedSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListTemplateCategoryRecommendedSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListTemplateCategoryRecommendedSettingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTemplateCategoryRecommendedSettingsOperationOptions() ListTemplateCategoryRecommendedSettingsOperationOptions {
	return ListTemplateCategoryRecommendedSettingsOperationOptions{}
}

func (o ListTemplateCategoryRecommendedSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTemplateCategoryRecommendedSettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListTemplateCategoryRecommendedSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTemplateCategoryRecommendedSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateCategoryRecommendedSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateCategoryRecommendedSettings - Get recommendedSettings from deviceManagement. The settings this category
// contains
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettings(ctx context.Context, id beta.DeviceManagementTemplateIdCategoryId, options ListTemplateCategoryRecommendedSettingsOperationOptions) (result ListTemplateCategoryRecommendedSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTemplateCategoryRecommendedSettingsCustomPager{},
		Path:          fmt.Sprintf("%s/recommendedSettings", id.ID()),
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

// ListTemplateCategoryRecommendedSettingsComplete retrieves all the results into a single object
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettingsComplete(ctx context.Context, id beta.DeviceManagementTemplateIdCategoryId, options ListTemplateCategoryRecommendedSettingsOperationOptions) (ListTemplateCategoryRecommendedSettingsCompleteResult, error) {
	return c.ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate(ctx, id, options, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateCategoryRecommendedSettingClient) ListTemplateCategoryRecommendedSettingsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementTemplateIdCategoryId, options ListTemplateCategoryRecommendedSettingsOperationOptions, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListTemplateCategoryRecommendedSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListTemplateCategoryRecommendedSettings(ctx, id, options)
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

	result = ListTemplateCategoryRecommendedSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
