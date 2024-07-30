package calendargroupcalendarevent

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

type ListCalendarGroupCalendarEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Event
}

type ListCalendarGroupCalendarEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Event
}

type ListCalendarGroupCalendarEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarEvents ...
func (c CalendarGroupCalendarEventClient) ListCalendarGroupCalendarEvents(ctx context.Context, id MeCalendarGroupIdCalendarId) (result ListCalendarGroupCalendarEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarEventsCustomPager{},
		Path:       fmt.Sprintf("%s/events", id.ID()),
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

// ListCalendarGroupCalendarEventsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarEventClient) ListCalendarGroupCalendarEventsComplete(ctx context.Context, id MeCalendarGroupIdCalendarId) (ListCalendarGroupCalendarEventsCompleteResult, error) {
	return c.ListCalendarGroupCalendarEventsCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarGroupCalendarEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarEventClient) ListCalendarGroupCalendarEventsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupIdCalendarId, predicate EventOperationPredicate) (result ListCalendarGroupCalendarEventsCompleteResult, err error) {
	items := make([]stable.Event, 0)

	resp, err := c.ListCalendarGroupCalendarEvents(ctx, id)
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

	result = ListCalendarGroupCalendarEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
