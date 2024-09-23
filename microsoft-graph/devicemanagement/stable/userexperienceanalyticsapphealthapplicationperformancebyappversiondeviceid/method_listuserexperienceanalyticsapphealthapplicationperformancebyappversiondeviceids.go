package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

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

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions() ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions {
	return ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions{}
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds - Get
// userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId from deviceManagement. User experience
// analytics appHealth Application Performance by App Version Device Id
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId",
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
		Values *[]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsComplete(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions) (ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate{})
}

// ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient) ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsOperationOptions, predicate UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdOperationPredicate) (result ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId, 0)

	resp, err := c.ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIds(ctx, options)
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

	result = ListUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
