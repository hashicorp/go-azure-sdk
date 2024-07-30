package calendarviewinstanceexceptionoccurrenceattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarViewInstanceExceptionOccurrenceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewInstanceExceptionOccurrenceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewInstanceExceptionOccurrenceAttachments ...
func (c CalendarViewInstanceExceptionOccurrenceAttachmentClient) ListCalendarViewInstanceExceptionOccurrenceAttachments(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId) (result ListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarViewInstanceExceptionOccurrenceAttachmentsCustomPager{},
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
		Values *[]beta.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarViewInstanceExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c CalendarViewInstanceExceptionOccurrenceAttachmentClient) ListCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId) (ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewInstanceExceptionOccurrenceAttachmentClient) ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId, predicate AttachmentOperationPredicate) (result ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListCalendarViewInstanceExceptionOccurrenceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
