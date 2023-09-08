package usercalendareventinstanceexceptionoccurrenceextension

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

type ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensions ...
func (c UserCalendarEventInstanceExceptionOccurrenceExtensionClient) ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensions(ctx context.Context, id UserCalendarEventInstanceExceptionOccurrenceId) (result ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsOperationResponse, err error) {
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

// ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsComplete retrieves all the results into a single object
func (c UserCalendarEventInstanceExceptionOccurrenceExtensionClient) ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsComplete(ctx context.Context, id UserCalendarEventInstanceExceptionOccurrenceId) (ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteResult, error) {
	return c.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarEventInstanceExceptionOccurrenceExtensionClient) ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id UserCalendarEventInstanceExceptionOccurrenceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensions(ctx, id)
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

	result = ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
