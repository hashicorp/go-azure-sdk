package calendarcalendarviewattachment

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

type ListCalendarCalendarViewAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarCalendarViewAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarCalendarViewAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViewAttachments ...
func (c CalendarCalendarViewAttachmentClient) ListCalendarCalendarViewAttachments(ctx context.Context, id MeCalendarIdCalendarViewId) (result ListCalendarCalendarViewAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewAttachmentsCustomPager{},
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

// ListCalendarCalendarViewAttachmentsComplete retrieves all the results into a single object
func (c CalendarCalendarViewAttachmentClient) ListCalendarCalendarViewAttachmentsComplete(ctx context.Context, id MeCalendarIdCalendarViewId) (ListCalendarCalendarViewAttachmentsCompleteResult, error) {
	return c.ListCalendarCalendarViewAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarCalendarViewAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewAttachmentClient) ListCalendarCalendarViewAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarIdCalendarViewId, predicate AttachmentOperationPredicate) (result ListCalendarCalendarViewAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarCalendarViewAttachments(ctx, id)
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

	result = ListCalendarCalendarViewAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
