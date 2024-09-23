package settingwindow

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

type ListSettingWindowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsSetting
}

type ListSettingWindowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsSetting
}

type ListSettingWindowsOperationOptions struct {
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

func DefaultListSettingWindowsOperationOptions() ListSettingWindowsOperationOptions {
	return ListSettingWindowsOperationOptions{}
}

func (o ListSettingWindowsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSettingWindowsOperationOptions) ToOData() *odata.Query {
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

func (o ListSettingWindowsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSettingWindowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSettingWindowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSettingWindows - List Windows settings. Get a list of windowsSetting objects and their properties for the signed
// in user.
func (c SettingWindowClient) ListSettingWindows(ctx context.Context, options ListSettingWindowsOperationOptions) (result ListSettingWindowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSettingWindowsCustomPager{},
		Path:          "/me/settings/windows",
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
		Values *[]beta.WindowsSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSettingWindowsComplete retrieves all the results into a single object
func (c SettingWindowClient) ListSettingWindowsComplete(ctx context.Context, options ListSettingWindowsOperationOptions) (ListSettingWindowsCompleteResult, error) {
	return c.ListSettingWindowsCompleteMatchingPredicate(ctx, options, WindowsSettingOperationPredicate{})
}

// ListSettingWindowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SettingWindowClient) ListSettingWindowsCompleteMatchingPredicate(ctx context.Context, options ListSettingWindowsOperationOptions, predicate WindowsSettingOperationPredicate) (result ListSettingWindowsCompleteResult, err error) {
	items := make([]beta.WindowsSetting, 0)

	resp, err := c.ListSettingWindows(ctx, options)
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

	result = ListSettingWindowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
