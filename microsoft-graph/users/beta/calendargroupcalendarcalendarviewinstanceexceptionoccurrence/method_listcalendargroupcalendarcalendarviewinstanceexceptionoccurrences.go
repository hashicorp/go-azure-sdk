package calendargroupcalendarcalendarviewinstanceexceptionoccurrence

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

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrences ...
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrences(ctx context.Context, id UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId) (result ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCustomPager{},
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

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesComplete retrieves all the results into a single object
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesComplete(ctx context.Context, id UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId) (ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteResult, error) {
	return c.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId, predicate EventOperationPredicate) (result ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrences(ctx, id)
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

	result = ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
