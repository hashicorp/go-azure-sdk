package mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment

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

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments ...
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attachments", id.ID()),
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
		Values *[]models.AttachmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId) (ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarCalendarViewExceptionOccurrenceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
