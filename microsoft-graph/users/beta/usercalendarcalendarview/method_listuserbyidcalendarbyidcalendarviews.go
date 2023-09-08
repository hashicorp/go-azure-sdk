package usercalendarcalendarview

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

type ListUserByIdCalendarByIdCalendarViewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarByIdCalendarViewsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarByIdCalendarViews ...
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarByIdCalendarViews(ctx context.Context, id UserCalendarId) (result ListUserByIdCalendarByIdCalendarViewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendarView", id.ID()),
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

// ListUserByIdCalendarByIdCalendarViewsComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarByIdCalendarViewsComplete(ctx context.Context, id UserCalendarId) (ListUserByIdCalendarByIdCalendarViewsCompleteResult, error) {
	return c.ListUserByIdCalendarByIdCalendarViewsCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdCalendarViewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewClient) ListUserByIdCalendarByIdCalendarViewsCompleteMatchingPredicate(ctx context.Context, id UserCalendarId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdCalendarViewsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdCalendarViews(ctx, id)
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

	result = ListUserByIdCalendarByIdCalendarViewsCompleteResult{
		Items: items,
	}
	return
}
