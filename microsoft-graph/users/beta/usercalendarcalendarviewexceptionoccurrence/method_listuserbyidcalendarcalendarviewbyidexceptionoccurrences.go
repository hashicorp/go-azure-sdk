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

type ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarCalendarViewByIdExceptionOccurrences ...
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarCalendarViewByIdExceptionOccurrences(ctx context.Context, id UserCalendarCalendarViewId) (result ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesComplete(ctx context.Context, id UserCalendarCalendarViewId) (ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewExceptionOccurrenceClient) ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id UserCalendarCalendarViewId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarCalendarViewByIdExceptionOccurrences(ctx, id)
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

	result = ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
