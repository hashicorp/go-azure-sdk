package calendarview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarViewsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarViewsOperationOptions struct {
	Count         *bool
	EndDateTime   *string
	Expand        *odata.Expand
	Filter        *string
	Metadata      *odata.Metadata
	OrderBy       *odata.OrderBy
	RetryFunc     client.RequestRetryFunc
	Search        *string
	Select        *[]string
	Skip          *int64
	StartDateTime *string
	Top           *int64
}

func DefaultListCalendarViewsOperationOptions() ListCalendarViewsOperationOptions {
	return ListCalendarViewsOperationOptions{}
}

func (o ListCalendarViewsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarViewsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarViewsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDateTime != nil {
		out.Append("endDateTime", fmt.Sprintf("%v", *o.EndDateTime))
	}
	if o.StartDateTime != nil {
		out.Append("startDateTime", fmt.Sprintf("%v", *o.StartDateTime))
	}
	return &out
}

type ListCalendarViewsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViews - List calendarView. Get the occurrences, exceptions, and single instances of events in a calendar
// view defined by a time range, from the default calendar of a group.
func (c CalendarViewClient) ListCalendarViews(ctx context.Context, id beta.GroupId, options ListCalendarViewsOperationOptions) (result ListCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarViewsCustomPager{},
		Path:          fmt.Sprintf("%s/calendarView", id.ID()),
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
		Values *[]beta.Event `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarViewsComplete retrieves all the results into a single object
func (c CalendarViewClient) ListCalendarViewsComplete(ctx context.Context, id beta.GroupId, options ListCalendarViewsOperationOptions) (ListCalendarViewsCompleteResult, error) {
	return c.ListCalendarViewsCompleteMatchingPredicate(ctx, id, options, EventOperationPredicate{})
}

// ListCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewClient) ListCalendarViewsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListCalendarViewsOperationOptions, predicate EventOperationPredicate) (result ListCalendarViewsCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarViews(ctx, id, options)
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

	result = ListCalendarViewsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
