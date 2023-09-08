package usercalendarcalendarview

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

type ListUserByIdCalendarCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarCalendarViewsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarCalendarViews ...
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarCalendarViews(ctx context.Context, id UserId) (result ListUserByIdCalendarCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.EventCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdCalendarCalendarViewsComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarCalendarViewsComplete(ctx context.Context, id UserId) (ListUserByIdCalendarCalendarViewsCompleteResult, error) {
	return c.ListUserByIdCalendarCalendarViewsCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarCalendarViewsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarCalendarViewsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarCalendarViews(ctx, id)
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

	result = ListUserByIdCalendarCalendarViewsCompleteResult{
		Items: items,
	}
	return
}
