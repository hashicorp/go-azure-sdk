package userexperienceanalyticsapphealthapplicationperformancebyosversion

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions() ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions {
	return ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions{}
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersions - Get
// userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion from deviceManagement. User experience analytics
// appHealth Application Performance by OS Version
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersions(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion",
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
		Values *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsComplete(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsOperationOptions, predicate UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersions(ctx, options)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
