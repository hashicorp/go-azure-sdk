package deponboardingsetting

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

type ListDepOnboardingSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DepOnboardingSetting
}

type ListDepOnboardingSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DepOnboardingSetting
}

type ListDepOnboardingSettingsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDepOnboardingSettingsOperationOptions() ListDepOnboardingSettingsOperationOptions {
	return ListDepOnboardingSettingsOperationOptions{}
}

func (o ListDepOnboardingSettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDepOnboardingSettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListDepOnboardingSettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDepOnboardingSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettings - Get depOnboardingSettings from deviceManagement. This collections of multiple DEP tokens
// per-tenant.
func (c DepOnboardingSettingClient) ListDepOnboardingSettings(ctx context.Context, options ListDepOnboardingSettingsOperationOptions) (result ListDepOnboardingSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDepOnboardingSettingsCustomPager{},
		Path:          "/deviceManagement/depOnboardingSettings",
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
		Values *[]beta.DepOnboardingSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingsComplete retrieves all the results into a single object
func (c DepOnboardingSettingClient) ListDepOnboardingSettingsComplete(ctx context.Context, options ListDepOnboardingSettingsOperationOptions) (ListDepOnboardingSettingsCompleteResult, error) {
	return c.ListDepOnboardingSettingsCompleteMatchingPredicate(ctx, options, DepOnboardingSettingOperationPredicate{})
}

// ListDepOnboardingSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingClient) ListDepOnboardingSettingsCompleteMatchingPredicate(ctx context.Context, options ListDepOnboardingSettingsOperationOptions, predicate DepOnboardingSettingOperationPredicate) (result ListDepOnboardingSettingsCompleteResult, err error) {
	items := make([]beta.DepOnboardingSetting, 0)

	resp, err := c.ListDepOnboardingSettings(ctx, options)
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

	result = ListDepOnboardingSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
