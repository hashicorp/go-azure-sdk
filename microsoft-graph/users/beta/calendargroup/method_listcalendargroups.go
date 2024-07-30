package calendargroup

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

type ListCalendarGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CalendarGroup
}

type ListCalendarGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CalendarGroup
}

type ListCalendarGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroups ...
func (c CalendarGroupClient) ListCalendarGroups(ctx context.Context, id UserId) (result ListCalendarGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupsCustomPager{},
		Path:       fmt.Sprintf("%s/calendarGroups", id.ID()),
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
		Values *[]beta.CalendarGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarGroupsComplete retrieves all the results into a single object
func (c CalendarGroupClient) ListCalendarGroupsComplete(ctx context.Context, id UserId) (ListCalendarGroupsCompleteResult, error) {
	return c.ListCalendarGroupsCompleteMatchingPredicate(ctx, id, CalendarGroupOperationPredicate{})
}

// ListCalendarGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupClient) ListCalendarGroupsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate CalendarGroupOperationPredicate) (result ListCalendarGroupsCompleteResult, err error) {
	items := make([]beta.CalendarGroup, 0)

	resp, err := c.ListCalendarGroups(ctx, id)
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

	result = ListCalendarGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
