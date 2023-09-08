package usercalendarcalendarviewexceptionoccurrence

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

type ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrences ...
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrences(ctx context.Context, id UserCalendarCalendarViewId) (result ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/exceptionOccurrences", id.ID()),
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

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesComplete(ctx context.Context, id UserCalendarCalendarViewId) (ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id UserCalendarCalendarViewId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrences(ctx, id)
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

	result = ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
