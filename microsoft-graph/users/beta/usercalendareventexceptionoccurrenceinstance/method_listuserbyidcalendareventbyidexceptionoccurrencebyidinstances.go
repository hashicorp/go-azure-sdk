package usercalendareventexceptionoccurrenceinstance

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

type ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstances ...
func (c UserCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstances(ctx context.Context, id UserCalendarEventExceptionOccurrenceId) (result ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesOperationResponse, err error) {
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

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesComplete retrieves all the results into a single object
func (c UserCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesComplete(ctx context.Context, id UserCalendarEventExceptionOccurrenceId) (ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteResult, error) {
	return c.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx context.Context, id UserCalendarEventExceptionOccurrenceId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstances(ctx, id)
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

	result = ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
