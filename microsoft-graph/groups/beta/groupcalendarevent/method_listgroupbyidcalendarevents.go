package groupcalendarevent

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

type ListGroupByIdCalendarEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.EventCollectionResponse
}

type ListGroupByIdCalendarEventsCompleteResult struct {
	Items []models.EventCollectionResponse
}

// ListGroupByIdCalendarEvents ...
func (c GroupCalendarEventClient) ListGroupByIdCalendarEvents(ctx context.Context, id GroupId) (result ListGroupByIdCalendarEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/calendar/events", id.ID()),
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

// ListGroupByIdCalendarEventsComplete retrieves all the results into a single object
func (c GroupCalendarEventClient) ListGroupByIdCalendarEventsComplete(ctx context.Context, id GroupId) (ListGroupByIdCalendarEventsCompleteResult, error) {
	return c.ListGroupByIdCalendarEventsCompleteMatchingPredicate(ctx, id, models.EventCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarEventClient) ListGroupByIdCalendarEventsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.EventCollectionResponseOperationPredicate) (result ListGroupByIdCalendarEventsCompleteResult, err error) {
	items := make([]models.EventCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarEvents(ctx, id)
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

	result = ListGroupByIdCalendarEventsCompleteResult{
		Items: items,
	}
	return
}
