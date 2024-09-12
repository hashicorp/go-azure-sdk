package calendargroupcalendarcalendarviewexceptionoccurrenceattachment

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

type ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions() ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions {
	return ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions{}
}

func (o ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarViewExceptionOccurrenceAttachments - Get attachments from me. The collection of
// FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only.
// Nullable.
func (c CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarGroupCalendarViewExceptionOccurrenceAttachments(ctx context.Context, id beta.MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, options ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions) (result ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCustomPager{},
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

// ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsComplete(ctx context.Context, id beta.MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, options ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions) (ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentOperationPredicate{})
}

// ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient) ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id beta.MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceId, options ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions, predicate AttachmentOperationPredicate) (result ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListCalendarGroupCalendarViewExceptionOccurrenceAttachments(ctx, id, options)
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

	result = ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
