package calendarviewexceptionoccurrence

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

type ListCalendarViewExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarViewExceptionOccurrencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarViewExceptionOccurrencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewExceptionOccurrencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewExceptionOccurrences ...
func (c CalendarViewExceptionOccurrenceClient) ListCalendarViewExceptionOccurrences(ctx context.Context, id MeCalendarViewId) (result ListCalendarViewExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarViewExceptionOccurrencesCustomPager{},
		Path:       fmt.Sprintf("%s/exceptionOccurrences", id.ID()),
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

// ListCalendarViewExceptionOccurrencesComplete retrieves all the results into a single object
func (c CalendarViewExceptionOccurrenceClient) ListCalendarViewExceptionOccurrencesComplete(ctx context.Context, id MeCalendarViewId) (ListCalendarViewExceptionOccurrencesCompleteResult, error) {
	return c.ListCalendarViewExceptionOccurrencesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarViewExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewExceptionOccurrenceClient) ListCalendarViewExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewId, predicate EventOperationPredicate) (result ListCalendarViewExceptionOccurrencesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarViewExceptionOccurrences(ctx, id)
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

	result = ListCalendarViewExceptionOccurrencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
