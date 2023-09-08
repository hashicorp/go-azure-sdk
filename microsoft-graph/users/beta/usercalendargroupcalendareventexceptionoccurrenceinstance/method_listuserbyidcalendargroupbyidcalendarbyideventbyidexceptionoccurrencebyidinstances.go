package usercalendargroupcalendareventexceptionoccurrenceinstance

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

type ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances ...
func (c UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx context.Context, id UserCalendarGroupCalendarEventExceptionOccurrenceId) (result ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesOperationResponse, err error) {
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

// ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete retrieves all the results into a single object
func (c UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx context.Context, id UserCalendarGroupCalendarEventExceptionOccurrenceId) (ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteResult, error) {
	return c.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient) ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx context.Context, id UserCalendarGroupCalendarEventExceptionOccurrenceId, predicate models.EventCollectionResponseOperationPredicate) (result ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)
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

	result = ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
