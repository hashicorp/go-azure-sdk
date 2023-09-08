package mecalendarcalendarviewexceptionoccurrence

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

type ListMeCalendarCalendarViewByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarCalendarViewByIdExceptionOccurrences ...
func (c MeCalendarCalendarViewExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdExceptionOccurrences(ctx context.Context, id MeCalendarCalendarViewId) (result ListMeCalendarCalendarViewByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListMeCalendarCalendarViewByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdExceptionOccurrencesComplete(ctx context.Context, id MeCalendarCalendarViewId) (ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarCalendarViewByIdExceptionOccurrences(ctx, id)
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

	result = ListMeCalendarCalendarViewByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
