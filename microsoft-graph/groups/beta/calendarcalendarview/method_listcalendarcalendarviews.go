package calendarcalendarview

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

type ListCalendarCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarCalendarViewsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarCalendarViewsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViews ...
func (c CalendarCalendarViewClient) ListCalendarCalendarViews(ctx context.Context, id GroupId) (result ListCalendarCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewsCustomPager{},
		Path:       fmt.Sprintf("%s/calendar/calendarView", id.ID()),
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

// ListCalendarCalendarViewsComplete retrieves all the results into a single object
func (c CalendarCalendarViewClient) ListCalendarCalendarViewsComplete(ctx context.Context, id GroupId) (ListCalendarCalendarViewsCompleteResult, error) {
	return c.ListCalendarCalendarViewsCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewClient) ListCalendarCalendarViewsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate EventOperationPredicate) (result ListCalendarCalendarViewsCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarCalendarViews(ctx, id)
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

	result = ListCalendarCalendarViewsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
