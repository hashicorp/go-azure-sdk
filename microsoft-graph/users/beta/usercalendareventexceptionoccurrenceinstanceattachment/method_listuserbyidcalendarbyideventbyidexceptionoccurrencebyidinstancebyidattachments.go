package usercalendareventexceptionoccurrenceinstanceattachment

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

type ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments ...
func (c UserCalendarEventExceptionOccurrenceInstanceAttachmentClient) ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx context.Context, id UserCalendarEventExceptionOccurrenceInstanceId) (result ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete retrieves all the results into a single object
func (c UserCalendarEventExceptionOccurrenceInstanceAttachmentClient) ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx context.Context, id UserCalendarEventExceptionOccurrenceInstanceId) (ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarEventExceptionOccurrenceInstanceAttachmentClient) ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserCalendarEventExceptionOccurrenceInstanceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)
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

	result = ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
