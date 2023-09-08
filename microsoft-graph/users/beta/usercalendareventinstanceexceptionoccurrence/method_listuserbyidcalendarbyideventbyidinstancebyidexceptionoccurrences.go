package usercalendareventinstanceexceptionoccurrence

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

type ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences ...
func (c UserCalendarEventInstanceExceptionOccurrenceClient) ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx context.Context, id UserCalendarEventInstanceId) (result ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c UserCalendarEventInstanceExceptionOccurrenceClient) ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx context.Context, id UserCalendarEventInstanceId) (ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarEventInstanceExceptionOccurrenceClient) ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id UserCalendarEventInstanceId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)
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

	result = ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
