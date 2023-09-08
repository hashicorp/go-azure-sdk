package mecalendargroupcalendareventinstanceexceptionoccurrenceattachment

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

type ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments ...
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx context.Context, id MeCalendarGroupCalendarEventInstanceExceptionOccurrenceId) (result ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx context.Context, id MeCalendarGroupCalendarEventInstanceExceptionOccurrenceId) (ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarEventInstanceExceptionOccurrenceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
