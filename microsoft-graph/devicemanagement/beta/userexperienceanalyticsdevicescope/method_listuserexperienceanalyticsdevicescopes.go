package userexperienceanalyticsdevicescope

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

type ListUserExperienceAnalyticsDeviceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsDeviceScope
}

type ListUserExperienceAnalyticsDeviceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsDeviceScope
}

type ListUserExperienceAnalyticsDeviceScopesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsDeviceScopesOperationOptions() ListUserExperienceAnalyticsDeviceScopesOperationOptions {
	return ListUserExperienceAnalyticsDeviceScopesOperationOptions{}
}

func (o ListUserExperienceAnalyticsDeviceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsDeviceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsDeviceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsDeviceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceScopes - Get userExperienceAnalyticsDeviceScopes from deviceManagement. The user
// experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint
// analytics reports.
func (c UserExperienceAnalyticsDeviceScopeClient) ListUserExperienceAnalyticsDeviceScopes(ctx context.Context, options ListUserExperienceAnalyticsDeviceScopesOperationOptions) (result ListUserExperienceAnalyticsDeviceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsDeviceScopesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceScopes",
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
		Values *[]beta.UserExperienceAnalyticsDeviceScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceScopesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceScopeClient) ListUserExperienceAnalyticsDeviceScopesComplete(ctx context.Context, options ListUserExperienceAnalyticsDeviceScopesOperationOptions) (ListUserExperienceAnalyticsDeviceScopesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceScopesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsDeviceScopeOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceScopeClient) ListUserExperienceAnalyticsDeviceScopesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsDeviceScopesOperationOptions, predicate UserExperienceAnalyticsDeviceScopeOperationPredicate) (result ListUserExperienceAnalyticsDeviceScopesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsDeviceScope, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceScopes(ctx, options)
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

	result = ListUserExperienceAnalyticsDeviceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
