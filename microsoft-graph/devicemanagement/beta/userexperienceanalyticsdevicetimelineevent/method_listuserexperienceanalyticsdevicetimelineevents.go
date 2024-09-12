package userexperienceanalyticsdevicetimelineevent

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

type ListUserExperienceAnalyticsDeviceTimelineEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsDeviceTimelineEvent
}

type ListUserExperienceAnalyticsDeviceTimelineEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsDeviceTimelineEvent
}

type ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions() ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions {
	return ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions{}
}

func (o ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsDeviceTimelineEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceTimelineEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceTimelineEvents - Get userExperienceAnalyticsDeviceTimelineEvent from
// deviceManagement. The user experience analytics device events entity contains NRT device timeline event details.
func (c UserExperienceAnalyticsDeviceTimelineEventClient) ListUserExperienceAnalyticsDeviceTimelineEvents(ctx context.Context, options ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions) (result ListUserExperienceAnalyticsDeviceTimelineEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsDeviceTimelineEventsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent",
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
		Values *[]beta.UserExperienceAnalyticsDeviceTimelineEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceTimelineEventsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceTimelineEventClient) ListUserExperienceAnalyticsDeviceTimelineEventsComplete(ctx context.Context, options ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions) (ListUserExperienceAnalyticsDeviceTimelineEventsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceTimelineEventsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsDeviceTimelineEventOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceTimelineEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceTimelineEventClient) ListUserExperienceAnalyticsDeviceTimelineEventsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsDeviceTimelineEventsOperationOptions, predicate UserExperienceAnalyticsDeviceTimelineEventOperationPredicate) (result ListUserExperienceAnalyticsDeviceTimelineEventsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsDeviceTimelineEvent, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceTimelineEvents(ctx, options)
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

	result = ListUserExperienceAnalyticsDeviceTimelineEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
