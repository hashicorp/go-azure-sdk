package mecalendarviewexceptionoccurrenceinstance

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

type ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarViewByIdExceptionOccurrenceByIdInstances ...
func (c MeCalendarViewExceptionOccurrenceInstanceClient) ListMeCalendarViewByIdExceptionOccurrenceByIdInstances(ctx context.Context, id MeCalendarViewExceptionOccurrenceId) (result ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesOperationResponse, err error) {
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

// ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesComplete retrieves all the results into a single object
func (c MeCalendarViewExceptionOccurrenceInstanceClient) ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx context.Context, id MeCalendarViewExceptionOccurrenceId) (ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult, error) {
	return c.ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarViewExceptionOccurrenceInstanceClient) ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewExceptionOccurrenceId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)
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

	result = ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesCompleteResult{
		Items: items,
	}
	return
}
