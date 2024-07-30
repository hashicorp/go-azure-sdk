package calendarcalendarviewexceptionoccurrenceinstanceattachment

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

type ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViewExceptionOccurrenceInstanceAttachments ...
func (c CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceInstanceAttachments(ctx context.Context, id MeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceId) (result ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCustomPager{},
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

// ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsComplete retrieves all the results into a single object
func (c CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx context.Context, id MeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceId) (ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteResult, error) {
	return c.ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceId, predicate AttachmentOperationPredicate) (result ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id)
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

	result = ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
