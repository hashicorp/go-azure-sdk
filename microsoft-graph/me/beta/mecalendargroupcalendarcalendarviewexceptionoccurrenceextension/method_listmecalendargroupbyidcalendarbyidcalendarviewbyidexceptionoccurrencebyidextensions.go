package mecalendargroupcalendarcalendarviewexceptionoccurrenceextension

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

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensions ...
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensions(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsComplete(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) (ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensions(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
