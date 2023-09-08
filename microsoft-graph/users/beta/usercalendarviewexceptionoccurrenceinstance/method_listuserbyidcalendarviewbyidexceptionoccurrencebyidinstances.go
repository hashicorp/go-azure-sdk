package usercalendarviewexceptionoccurrenceinstance

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

type ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstances ...
func (c UserCalendarViewExceptionOccurrenceInstanceClient) ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstances(ctx context.Context, id UserCalendarViewExceptionOccurrenceId) (result ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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

// ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesComplete retrieves all the results into a single object
func (c UserCalendarViewExceptionOccurrenceInstanceClient) ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx context.Context, id UserCalendarViewExceptionOccurrenceId) (ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult, error) {
	return c.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarViewExceptionOccurrenceInstanceClient) ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx context.Context, id UserCalendarViewExceptionOccurrenceId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)
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

	result = ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
