package usercalendarviewinstanceexceptionoccurrence

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

type ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrences ...
func (c UserCalendarViewInstanceExceptionOccurrenceClient) ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx context.Context, id UserCalendarViewInstanceId) (result ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c UserCalendarViewInstanceExceptionOccurrenceClient) ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx context.Context, id UserCalendarViewInstanceId) (ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarViewInstanceExceptionOccurrenceClient) ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id UserCalendarViewInstanceId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)
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

	result = ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
