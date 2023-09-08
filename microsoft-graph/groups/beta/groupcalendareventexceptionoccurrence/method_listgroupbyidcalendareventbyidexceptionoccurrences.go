package groupcalendareventexceptionoccurrence

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

type ListGroupByIdCalendarEventByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListGroupByIdCalendarEventByIdExceptionOccurrences ...
func (c GroupCalendarEventExceptionOccurrenceClient) ListGroupByIdCalendarEventByIdExceptionOccurrences(ctx context.Context, id GroupCalendarEventId) (result ListGroupByIdCalendarEventByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListGroupByIdCalendarEventByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c GroupCalendarEventExceptionOccurrenceClient) ListGroupByIdCalendarEventByIdExceptionOccurrencesComplete(ctx context.Context, id GroupCalendarEventId) (ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarEventExceptionOccurrenceClient) ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id GroupCalendarEventId, predicate models.EventCollectionResponseOperationPredicate) (result ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarEventByIdExceptionOccurrences(ctx, id)
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

	result = ListGroupByIdCalendarEventByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
