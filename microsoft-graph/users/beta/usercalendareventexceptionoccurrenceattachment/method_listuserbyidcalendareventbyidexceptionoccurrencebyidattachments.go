package usercalendareventexceptionoccurrenceattachment

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

type ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachments ...
func (c UserCalendarEventExceptionOccurrenceAttachmentClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachments(ctx context.Context, id UserCalendarEventExceptionOccurrenceId) (result ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsComplete retrieves all the results into a single object
func (c UserCalendarEventExceptionOccurrenceAttachmentClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx context.Context, id UserCalendarEventExceptionOccurrenceId) (ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarEventExceptionOccurrenceAttachmentClient) ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserCalendarEventExceptionOccurrenceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachments(ctx, id)
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

	result = ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
