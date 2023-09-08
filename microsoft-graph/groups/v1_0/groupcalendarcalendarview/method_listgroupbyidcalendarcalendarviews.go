package groupcalendarcalendarview

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

type ListGroupByIdCalendarCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListGroupByIdCalendarCalendarViewsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListGroupByIdCalendarCalendarViews ...
func (c GroupCalendarCalendarViewClient) ListGroupByIdCalendarCalendarViews(ctx context.Context, id GroupId) (result ListGroupByIdCalendarCalendarViewsOperationResponse, err error) {
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

// ListGroupByIdCalendarCalendarViewsComplete retrieves all the results into a single object
func (c GroupCalendarCalendarViewClient) ListGroupByIdCalendarCalendarViewsComplete(ctx context.Context, id GroupId) (ListGroupByIdCalendarCalendarViewsCompleteResult, error) {
	return c.ListGroupByIdCalendarCalendarViewsCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarCalendarViewClient) ListGroupByIdCalendarCalendarViewsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.EventCollectionResponseOperationPredicate) (result ListGroupByIdCalendarCalendarViewsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarCalendarViews(ctx, id)
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

	result = ListGroupByIdCalendarCalendarViewsCompleteResult{
		Items: items,
	}
	return
}
