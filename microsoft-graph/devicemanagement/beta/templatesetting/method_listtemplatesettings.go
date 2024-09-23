package templatesetting

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

type ListTemplateSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingInstance
}

type ListTemplateSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingInstance
}

type ListTemplateSettingsOperationOptions struct {
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

func DefaultListTemplateSettingsOperationOptions() ListTemplateSettingsOperationOptions {
	return ListTemplateSettingsOperationOptions{}
}

func (o ListTemplateSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTemplateSettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListTemplateSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTemplateSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateSettings - Get settings from deviceManagement. Collection of all settings this template has
func (c TemplateSettingClient) ListTemplateSettings(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateSettingsOperationOptions) (result ListTemplateSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTemplateSettingsCustomPager{},
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

// ListTemplateSettingsComplete retrieves all the results into a single object
func (c TemplateSettingClient) ListTemplateSettingsComplete(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateSettingsOperationOptions) (ListTemplateSettingsCompleteResult, error) {
	return c.ListTemplateSettingsCompleteMatchingPredicate(ctx, id, options, DeviceManagementSettingInstanceOperationPredicate{})
}

// ListTemplateSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateSettingClient) ListTemplateSettingsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateSettingsOperationOptions, predicate DeviceManagementSettingInstanceOperationPredicate) (result ListTemplateSettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingInstance, 0)

	resp, err := c.ListTemplateSettings(ctx, id, options)
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

	result = ListTemplateSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
