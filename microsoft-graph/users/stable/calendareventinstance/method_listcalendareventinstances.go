package calendareventinstance

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

type ListCalendarEventInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Event
}

type ListCalendarEventInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Event
}

type ListCalendarEventInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarEventInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarEventInstances ...
func (c CalendarEventInstanceClient) ListCalendarEventInstances(ctx context.Context, id UserIdCalendarIdEventId) (result ListCalendarEventInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarEventInstancesCustomPager{},
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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

// ListCalendarEventInstancesComplete retrieves all the results into a single object
func (c CalendarEventInstanceClient) ListCalendarEventInstancesComplete(ctx context.Context, id UserIdCalendarIdEventId) (ListCalendarEventInstancesCompleteResult, error) {
	return c.ListCalendarEventInstancesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarEventInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarEventInstanceClient) ListCalendarEventInstancesCompleteMatchingPredicate(ctx context.Context, id UserIdCalendarIdEventId, predicate EventOperationPredicate) (result ListCalendarEventInstancesCompleteResult, err error) {
	items := make([]stable.Event, 0)

	resp, err := c.ListCalendarEventInstances(ctx, id)
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

	result = ListCalendarEventInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
