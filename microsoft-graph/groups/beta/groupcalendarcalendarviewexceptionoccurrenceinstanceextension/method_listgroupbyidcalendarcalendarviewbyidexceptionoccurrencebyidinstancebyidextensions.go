package groupcalendarcalendarviewexceptionoccurrenceinstanceextension

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

type ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions ...
func (c GroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions(ctx context.Context, id GroupCalendarCalendarViewExceptionOccurrenceInstanceId) (result ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]models.ExtensionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsComplete retrieves all the results into a single object
func (c GroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsComplete(ctx context.Context, id GroupCalendarCalendarViewExceptionOccurrenceInstanceId) (ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult, error) {
	return c.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id GroupCalendarCalendarViewExceptionOccurrenceInstanceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions(ctx, id)
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

	result = ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
