package mecalendarcalendarviewinstanceexceptionoccurrence

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

type ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrences ...
func (c MeCalendarCalendarViewInstanceExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrences(ctx context.Context, id MeCalendarCalendarViewInstanceId) (result ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewInstanceExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx context.Context, id MeCalendarCalendarViewInstanceId) (ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewInstanceExceptionOccurrenceClient) ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewInstanceId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)
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

	result = ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
