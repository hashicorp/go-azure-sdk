package calendarview

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

type ListCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Event
}

type ListCalendarViewsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Event
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

// ListCalendarViews ...
func (c CalendarViewClient) ListCalendarViews(ctx context.Context) (result ListCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarViewsCustomPager{},
		Path:       "/me/calendarView",
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

// ListCalendarViewsComplete retrieves all the results into a single object
func (c CalendarViewClient) ListCalendarViewsComplete(ctx context.Context) (ListCalendarViewsCompleteResult, error) {
	return c.ListCalendarViewsCompleteMatchingPredicate(ctx, EventOperationPredicate{})
}

// ListCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewClient) ListCalendarViewsCompleteMatchingPredicate(ctx context.Context, predicate EventOperationPredicate) (result ListCalendarViewsCompleteResult, err error) {
	items := make([]stable.Event, 0)

	resp, err := c.ListCalendarViews(ctx)
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
