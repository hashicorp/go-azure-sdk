package usercalendarcalendarviewexceptionoccurrenceattachment

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

type ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments ...
func (c UserCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx context.Context, id UserCalendarCalendarViewExceptionOccurrenceId) (result ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx context.Context, id UserCalendarCalendarViewExceptionOccurrenceId) (ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserCalendarCalendarViewExceptionOccurrenceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)
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

	result = ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
