package userexperienceanalyticsapphealthapplicationperformancebyappversiondetail

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions() ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions {
	return ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions{}
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails - Get
// userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails from deviceManagement. User experience
// analytics appHealth Application Performance by App Version details
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails",
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
		Values *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsComplete(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsOperationOptions, predicate UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(ctx, options)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
