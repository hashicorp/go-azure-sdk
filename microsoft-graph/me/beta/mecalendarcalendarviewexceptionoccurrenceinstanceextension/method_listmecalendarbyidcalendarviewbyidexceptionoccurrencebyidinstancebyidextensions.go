package mecalendarcalendarviewexceptionoccurrenceinstanceextension

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

type ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions ...
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId) (result ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsOperationResponse, err error) {
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

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsComplete(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId) (ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult, error) {
	return c.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensions(ctx, id)
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

	result = ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
