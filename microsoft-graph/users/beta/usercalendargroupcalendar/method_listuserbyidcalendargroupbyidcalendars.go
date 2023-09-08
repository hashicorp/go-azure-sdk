package usercalendargroupcalendar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdCalendarGroupByIdCalendarsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarCollectionResponse
}

type ListUserByIdCalendarGroupByIdCalendarsCompleteResult struct {
	Items []models.CalendarCollectionResponse
}

// ListUserByIdCalendarGroupByIdCalendars ...
func (c UserCalendarGroupCalendarClient) ListUserByIdCalendarGroupByIdCalendars(ctx context.Context, id UserCalendarGroupId) (result ListUserByIdCalendarGroupByIdCalendarsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendars", id.ID()),
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
		Values *[]models.CalendarCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCalendarGroupByIdCalendarsComplete retrieves all the results into a single object
func (c UserCalendarGroupCalendarClient) ListUserByIdCalendarGroupByIdCalendarsComplete(ctx context.Context, id UserCalendarGroupId) (ListUserByIdCalendarGroupByIdCalendarsCompleteResult, error) {
	return c.ListUserByIdCalendarGroupByIdCalendarsCompleteMatchingPredicate(ctx, id, models.CalendarCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarGroupByIdCalendarsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarGroupCalendarClient) ListUserByIdCalendarGroupByIdCalendarsCompleteMatchingPredicate(ctx context.Context, id UserCalendarGroupId, predicate models.CalendarCollectionResponseOperationPredicate) (result ListUserByIdCalendarGroupByIdCalendarsCompleteResult, err error) {
	items := make([]models.CalendarCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarGroupByIdCalendars(ctx, id)
	if err != nil {
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

	result = ListUserByIdCalendarGroupByIdCalendarsCompleteResult{
		Items: items,
	}
	return
}
