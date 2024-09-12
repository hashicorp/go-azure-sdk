package calendareventinstanceexceptionoccurrenceattachment

import (
	"context"
	"encoding/json"
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

type ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions() ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions {
	return ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions{}
}

func (o ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListCalendarEventInstanceExceptionOccurrenceAttachments - Get attachments from groups. The collection of
// FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only.
// Nullable.
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx context.Context, id beta.GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId, options ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions) (result ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarEventInstanceExceptionOccurrenceAttachmentsCustomPager{},
		Path:          fmt.Sprintf("%s/attachments", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.Attachment, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalAttachmentImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.Attachment (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx context.Context, id beta.GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId, options ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions) (ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentOperationPredicate{})
}

// ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarEventInstanceExceptionOccurrenceAttachmentClient) ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdCalendarEventIdInstanceIdExceptionOccurrenceId, options ListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions, predicate AttachmentOperationPredicate) (result ListCalendarEventInstanceExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id, options)
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
