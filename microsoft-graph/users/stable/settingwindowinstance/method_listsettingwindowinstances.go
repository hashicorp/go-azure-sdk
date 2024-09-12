package settingwindowinstance

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

type ListSettingWindowInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.WindowsSettingInstance
}

type ListSettingWindowInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.WindowsSettingInstance
}

type ListSettingWindowInstancesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListSettingWindowInstancesOperationOptions() ListSettingWindowInstancesOperationOptions {
	return ListSettingWindowInstancesOperationOptions{}
}

func (o ListSettingWindowInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSettingWindowInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListSettingWindowInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSettingWindowInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSettingWindowInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSettingWindowInstances - Get instances from users. A collection of setting values for a given windowsSetting.
func (c SettingWindowInstanceClient) ListSettingWindowInstances(ctx context.Context, id stable.UserIdSettingWindowId, options ListSettingWindowInstancesOperationOptions) (result ListSettingWindowInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSettingWindowInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]stable.WindowsSettingInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSettingWindowInstancesComplete retrieves all the results into a single object
func (c SettingWindowInstanceClient) ListSettingWindowInstancesComplete(ctx context.Context, id stable.UserIdSettingWindowId, options ListSettingWindowInstancesOperationOptions) (ListSettingWindowInstancesCompleteResult, error) {
	return c.ListSettingWindowInstancesCompleteMatchingPredicate(ctx, id, options, WindowsSettingInstanceOperationPredicate{})
}

// ListSettingWindowInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SettingWindowInstanceClient) ListSettingWindowInstancesCompleteMatchingPredicate(ctx context.Context, id stable.UserIdSettingWindowId, options ListSettingWindowInstancesOperationOptions, predicate WindowsSettingInstanceOperationPredicate) (result ListSettingWindowInstancesCompleteResult, err error) {
	items := make([]stable.WindowsSettingInstance, 0)

	resp, err := c.ListSettingWindowInstances(ctx, id, options)
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

	result = ListSettingWindowInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
