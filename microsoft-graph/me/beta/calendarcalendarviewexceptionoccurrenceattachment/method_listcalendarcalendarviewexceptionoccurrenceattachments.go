package calendarcalendarviewexceptionoccurrenceattachment

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

type ListCalendarCalendarViewExceptionOccurrenceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarCalendarViewExceptionOccurrenceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewExceptionOccurrenceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViewExceptionOccurrenceAttachments ...
func (c CalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceAttachments(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceId) (result ListCalendarCalendarViewExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewExceptionOccurrenceAttachmentsCustomPager{},
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

// ListCalendarCalendarViewExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c CalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceAttachmentsComplete(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceId) (ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceId, predicate AttachmentOperationPredicate) (result ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarCalendarViewExceptionOccurrenceAttachments(ctx, id)
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

	result = ListCalendarCalendarViewExceptionOccurrenceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
