package calendareventinstanceexceptionoccurrenceattachment

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

type ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarEventInstanceExceptionOccurrenceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarEventInstanceExceptionOccurrenceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarEventInstanceExceptionOccurrenceAttachments ...
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx context.Context, id GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId) (result ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarEventInstanceExceptionOccurrenceAttachmentsCustomPager{},
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

// ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx context.Context, id GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId) (ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId, predicate AttachmentOperationPredicate) (result ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id)
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

	result = ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
