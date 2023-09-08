package mecalendarviewexceptionoccurrence

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

type ListMeCalendarViewByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarViewByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarViewByIdExceptionOccurrences ...
func (c MeCalendarViewExceptionOccurrenceClient) ListMeCalendarViewByIdExceptionOccurrences(ctx context.Context, id MeCalendarViewId) (result ListMeCalendarViewByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListMeCalendarViewByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c MeCalendarViewExceptionOccurrenceClient) ListMeCalendarViewByIdExceptionOccurrencesComplete(ctx context.Context, id MeCalendarViewId) (ListMeCalendarViewByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListMeCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarViewExceptionOccurrenceClient) ListMeCalendarViewByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarViewByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarViewByIdExceptionOccurrences(ctx, id)
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

	result = ListMeCalendarViewByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
