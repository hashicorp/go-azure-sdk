package calendarcalendarpermission

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

type ListCalendarCalendarPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CalendarPermission
}

type ListCalendarCalendarPermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CalendarPermission
}

type ListCalendarCalendarPermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarPermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarPermissions ...
func (c CalendarCalendarPermissionClient) ListCalendarCalendarPermissions(ctx context.Context, id UserId) (result ListCalendarCalendarPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarPermissionsCustomPager{},
		Path:       fmt.Sprintf("%s/calendar/calendarPermissions", id.ID()),
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
		Values *[]stable.CalendarPermission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarCalendarPermissionsComplete retrieves all the results into a single object
func (c CalendarCalendarPermissionClient) ListCalendarCalendarPermissionsComplete(ctx context.Context, id UserId) (ListCalendarCalendarPermissionsCompleteResult, error) {
	return c.ListCalendarCalendarPermissionsCompleteMatchingPredicate(ctx, id, CalendarPermissionOperationPredicate{})
}

// ListCalendarCalendarPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarPermissionClient) ListCalendarCalendarPermissionsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate CalendarPermissionOperationPredicate) (result ListCalendarCalendarPermissionsCompleteResult, err error) {
	items := make([]stable.CalendarPermission, 0)

	resp, err := c.ListCalendarCalendarPermissions(ctx, id)
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

	result = ListCalendarCalendarPermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
