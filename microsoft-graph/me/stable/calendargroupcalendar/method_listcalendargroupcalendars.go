package calendargroupcalendar

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

type ListCalendarGroupCalendarsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Calendar
}

type ListCalendarGroupCalendarsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Calendar
}

type ListCalendarGroupCalendarsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarGroupCalendarsOperationOptions() ListCalendarGroupCalendarsOperationOptions {
	return ListCalendarGroupCalendarsOperationOptions{}
}

func (o ListCalendarGroupCalendarsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarGroupCalendarsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarGroupCalendarsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarGroupCalendarsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendars - List calendars. Retrieve a list of calendars belonging to a calendar group.
func (c CalendarGroupCalendarClient) ListCalendarGroupCalendars(ctx context.Context, id stable.MeCalendarGroupId, options ListCalendarGroupCalendarsOperationOptions) (result ListCalendarGroupCalendarsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarGroupCalendarsCustomPager{},
		Path:          fmt.Sprintf("%s/calendars", id.ID()),
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
		Values *[]stable.Calendar `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarGroupCalendarsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarClient) ListCalendarGroupCalendarsComplete(ctx context.Context, id stable.MeCalendarGroupId, options ListCalendarGroupCalendarsOperationOptions) (ListCalendarGroupCalendarsCompleteResult, error) {
	return c.ListCalendarGroupCalendarsCompleteMatchingPredicate(ctx, id, options, CalendarOperationPredicate{})
}

// ListCalendarGroupCalendarsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarClient) ListCalendarGroupCalendarsCompleteMatchingPredicate(ctx context.Context, id stable.MeCalendarGroupId, options ListCalendarGroupCalendarsOperationOptions, predicate CalendarOperationPredicate) (result ListCalendarGroupCalendarsCompleteResult, err error) {
	items := make([]stable.Calendar, 0)

	resp, err := c.ListCalendarGroupCalendars(ctx, id, options)
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

	result = ListCalendarGroupCalendarsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
