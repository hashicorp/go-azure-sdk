package calendarviewinstance

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

type ListCalendarViewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Event
}

type ListCalendarViewInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Event
}

type ListCalendarViewInstancesOperationOptions struct {
	Count         *bool
	EndDateTime   *string
	Expand        *odata.Expand
	Filter        *string
	OrderBy       *odata.OrderBy
	Search        *string
	Select        *[]string
	Skip          *int64
	StartDateTime *string
	Top           *int64
}

func DefaultListCalendarViewInstancesOperationOptions() ListCalendarViewInstancesOperationOptions {
	return ListCalendarViewInstancesOperationOptions{}
}

func (o ListCalendarViewInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarViewInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarViewInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDateTime != nil {
		out.Append("endDateTime", fmt.Sprintf("%v", *o.EndDateTime))
	}
	if o.StartDateTime != nil {
		out.Append("startDateTime", fmt.Sprintf("%v", *o.StartDateTime))
	}
	return &out
}

type ListCalendarViewInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewInstances - Get instances from me. The occurrences of a recurring series, if the event is a series
// master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been
// modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only.
// Nullable.
func (c CalendarViewInstanceClient) ListCalendarViewInstances(ctx context.Context, id stable.MeCalendarViewId, options ListCalendarViewInstancesOperationOptions) (result ListCalendarViewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarViewInstancesCustomPager{},
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
		Values *[]stable.Event `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarViewInstancesComplete retrieves all the results into a single object
func (c CalendarViewInstanceClient) ListCalendarViewInstancesComplete(ctx context.Context, id stable.MeCalendarViewId, options ListCalendarViewInstancesOperationOptions) (ListCalendarViewInstancesCompleteResult, error) {
	return c.ListCalendarViewInstancesCompleteMatchingPredicate(ctx, id, options, EventOperationPredicate{})
}

// ListCalendarViewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewInstanceClient) ListCalendarViewInstancesCompleteMatchingPredicate(ctx context.Context, id stable.MeCalendarViewId, options ListCalendarViewInstancesOperationOptions, predicate EventOperationPredicate) (result ListCalendarViewInstancesCompleteResult, err error) {
	items := make([]stable.Event, 0)

	resp, err := c.ListCalendarViewInstances(ctx, id, options)
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

	result = ListCalendarViewInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
