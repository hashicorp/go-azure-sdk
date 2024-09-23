package calendarcalendarviewexceptionoccurrenceinstance

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

type ListCalendarViewExceptionOccurrenceInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarViewExceptionOccurrenceInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarViewExceptionOccurrenceInstancesOperationOptions struct {
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

func DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions() ListCalendarViewExceptionOccurrenceInstancesOperationOptions {
	return ListCalendarViewExceptionOccurrenceInstancesOperationOptions{}
}

func (o ListCalendarViewExceptionOccurrenceInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarViewExceptionOccurrenceInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarViewExceptionOccurrenceInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDateTime != nil {
		out.Append("endDateTime", fmt.Sprintf("%v", *o.EndDateTime))
	}
	if o.StartDateTime != nil {
		out.Append("startDateTime", fmt.Sprintf("%v", *o.StartDateTime))
	}
	return &out
}

type ListCalendarViewExceptionOccurrenceInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewExceptionOccurrenceInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewExceptionOccurrenceInstances - Get instances from groups. The occurrences of a recurring series, if
// the event is a series master. This property includes occurrences that are part of the recurrence pattern, and
// exceptions that have been modified, but doesn't include occurrences that have been canceled from the series.
// Navigation property. Read-only. Nullable.
func (c CalendarCalendarViewExceptionOccurrenceInstanceClient) ListCalendarViewExceptionOccurrenceInstances(ctx context.Context, id beta.GroupIdCalendarCalendarViewIdExceptionOccurrenceId, options ListCalendarViewExceptionOccurrenceInstancesOperationOptions) (result ListCalendarViewExceptionOccurrenceInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarViewExceptionOccurrenceInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/instances", id.ID()),
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

// ListCalendarViewExceptionOccurrenceInstancesComplete retrieves all the results into a single object
func (c CalendarCalendarViewExceptionOccurrenceInstanceClient) ListCalendarViewExceptionOccurrenceInstancesComplete(ctx context.Context, id beta.GroupIdCalendarCalendarViewIdExceptionOccurrenceId, options ListCalendarViewExceptionOccurrenceInstancesOperationOptions) (ListCalendarViewExceptionOccurrenceInstancesCompleteResult, error) {
	return c.ListCalendarViewExceptionOccurrenceInstancesCompleteMatchingPredicate(ctx, id, options, EventOperationPredicate{})
}

// ListCalendarViewExceptionOccurrenceInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewExceptionOccurrenceInstanceClient) ListCalendarViewExceptionOccurrenceInstancesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdCalendarCalendarViewIdExceptionOccurrenceId, options ListCalendarViewExceptionOccurrenceInstancesOperationOptions, predicate EventOperationPredicate) (result ListCalendarViewExceptionOccurrenceInstancesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarViewExceptionOccurrenceInstances(ctx, id, options)
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

	result = ListCalendarViewExceptionOccurrenceInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
