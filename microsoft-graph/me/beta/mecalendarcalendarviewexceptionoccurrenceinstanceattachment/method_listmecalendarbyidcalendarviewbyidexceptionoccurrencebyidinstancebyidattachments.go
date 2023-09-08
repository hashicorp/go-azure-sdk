package mecalendarcalendarviewexceptionoccurrenceinstanceattachment

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

type ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments ...
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId) (result ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsOperationResponse, err error) {
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

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId) (ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewExceptionOccurrenceInstanceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)
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

	result = ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
