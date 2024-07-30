package calendargroupcalendarcalendarview

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

type ListCalendarGroupCalendarCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListCalendarGroupCalendarCalendarViewsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListCalendarGroupCalendarCalendarViewsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarCalendarViewsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarCalendarViews ...
func (c CalendarGroupCalendarCalendarViewClient) ListCalendarGroupCalendarCalendarViews(ctx context.Context, id UserIdCalendarGroupIdCalendarId) (result ListCalendarGroupCalendarCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarCalendarViewsCustomPager{},
		Path:       fmt.Sprintf("%s/calendarView", id.ID()),
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

// ListCalendarGroupCalendarCalendarViewsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarCalendarViewClient) ListCalendarGroupCalendarCalendarViewsComplete(ctx context.Context, id UserIdCalendarGroupIdCalendarId) (ListCalendarGroupCalendarCalendarViewsCompleteResult, error) {
	return c.ListCalendarGroupCalendarCalendarViewsCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarGroupCalendarCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarCalendarViewClient) ListCalendarGroupCalendarCalendarViewsCompleteMatchingPredicate(ctx context.Context, id UserIdCalendarGroupIdCalendarId, predicate EventOperationPredicate) (result ListCalendarGroupCalendarCalendarViewsCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListCalendarGroupCalendarCalendarViews(ctx, id)
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

	result = ListCalendarGroupCalendarCalendarViewsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
