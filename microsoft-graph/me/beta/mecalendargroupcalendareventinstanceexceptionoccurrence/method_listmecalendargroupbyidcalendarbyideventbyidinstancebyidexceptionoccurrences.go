package mecalendargroupcalendareventinstanceexceptionoccurrence

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

type ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences ...
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx context.Context, id MeCalendarGroupCalendarEventInstanceId) (result ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx context.Context, id MeCalendarGroupCalendarEventInstanceId) (ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarEventInstanceId, predicate models.EventCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesCompleteResult{
		Items: items,
	}
	return
}
