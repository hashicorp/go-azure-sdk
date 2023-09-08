package mecalendargroupcalendar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeCalendarGroupByIdCalendarsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CalendarCollectionResponse
}

type ListMeCalendarGroupByIdCalendarsCompleteResult struct {
	Items []models.CalendarCollectionResponse
}

// ListMeCalendarGroupByIdCalendars ...
func (c MeCalendarGroupCalendarClient) ListMeCalendarGroupByIdCalendars(ctx context.Context, id MeCalendarGroupId) (result ListMeCalendarGroupByIdCalendarsOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarClient) ListMeCalendarGroupByIdCalendarsComplete(ctx context.Context, id MeCalendarGroupId) (ListMeCalendarGroupByIdCalendarsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarsCompleteMatchingPredicate(ctx, id, models.CalendarCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarClient) ListMeCalendarGroupByIdCalendarsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupId, predicate models.CalendarCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarsCompleteResult, err error) {
	items := make([]models.CalendarCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendars(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarsCompleteResult{
		Items: items,
	}
	return
}
