package calendargroupcalendareventinstanceexceptionoccurrence

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

type ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions() ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions {
	return ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions{}
}

func (o ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarEventInstanceExceptionOccurrences - Get exceptionOccurrences from users
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrences(ctx context.Context, id beta.UserIdCalendarGroupIdCalendarIdEventIdInstanceId, options ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions) (result ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCustomPager{},
		Path:          fmt.Sprintf("%s/exceptionOccurrences", id.ID()),
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

// ListCalendarGroupCalendarEventInstanceExceptionOccurrencesComplete retrieves all the results into a single object
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrencesComplete(ctx context.Context, id beta.UserIdCalendarGroupIdCalendarIdEventIdInstanceId, options ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions) (ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteResult, error) {
	return c.ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx, id, options, EventOperationPredicate{})
}

// ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id beta.UserIdCalendarGroupIdCalendarIdEventIdInstanceId, options ListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions, predicate EventOperationPredicate) (result ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarGroupCalendarEventInstanceExceptionOccurrences(ctx, id, options)
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

	result = ListCalendarGroupCalendarEventInstanceExceptionOccurrencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
