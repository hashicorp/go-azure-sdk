package calendarcalendarviewinstance

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

type ListCalendarCalendarViewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Event
}

type ListCalendarCalendarViewInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Event
}

type ListCalendarCalendarViewInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViewInstances ...
func (c CalendarCalendarViewInstanceClient) ListCalendarCalendarViewInstances(ctx context.Context, id MeCalendarCalendarViewId) (result ListCalendarCalendarViewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewInstancesCustomPager{},
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

// ListCalendarCalendarViewInstancesComplete retrieves all the results into a single object
func (c CalendarCalendarViewInstanceClient) ListCalendarCalendarViewInstancesComplete(ctx context.Context, id MeCalendarCalendarViewId) (ListCalendarCalendarViewInstancesCompleteResult, error) {
	return c.ListCalendarCalendarViewInstancesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListCalendarCalendarViewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewInstanceClient) ListCalendarCalendarViewInstancesCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewId, predicate EventOperationPredicate) (result ListCalendarCalendarViewInstancesCompleteResult, err error) {
	items := make([]stable.Event, 0)

	resp, err := c.ListCalendarCalendarViewInstances(ctx, id)
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

	result = ListCalendarCalendarViewInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
