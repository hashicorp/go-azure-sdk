package calendarviewattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCalendarViewAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Attachment
}

type ListCalendarViewAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Attachment
}

type ListCalendarViewAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewAttachments ...
func (c CalendarViewAttachmentClient) ListCalendarViewAttachments(ctx context.Context, id MeCalendarViewId) (result ListCalendarViewAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarViewAttachmentsCustomPager{},
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
		Values *[]stable.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarViewAttachmentsComplete retrieves all the results into a single object
func (c CalendarViewAttachmentClient) ListCalendarViewAttachmentsComplete(ctx context.Context, id MeCalendarViewId) (ListCalendarViewAttachmentsCompleteResult, error) {
	return c.ListCalendarViewAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListCalendarViewAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewAttachmentClient) ListCalendarViewAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewId, predicate AttachmentOperationPredicate) (result ListCalendarViewAttachmentsCompleteResult, err error) {
	items := make([]stable.Attachment, 0)

	resp, err := c.ListCalendarViewAttachments(ctx, id)
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

	result = ListCalendarViewAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
