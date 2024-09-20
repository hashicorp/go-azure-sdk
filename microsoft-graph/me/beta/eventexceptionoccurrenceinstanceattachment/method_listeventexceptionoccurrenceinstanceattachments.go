package eventexceptionoccurrenceinstanceattachment

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

type ListEventExceptionOccurrenceInstanceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListEventExceptionOccurrenceInstanceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListEventExceptionOccurrenceInstanceAttachmentsOperationOptions() ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions {
	return ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions{}
}

func (o ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEventExceptionOccurrenceInstanceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventExceptionOccurrenceInstanceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventExceptionOccurrenceInstanceAttachments - Get attachments from me. The collection of FileAttachment,
// ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (c EventExceptionOccurrenceInstanceAttachmentClient) ListEventExceptionOccurrenceInstanceAttachments(ctx context.Context, id beta.MeEventIdExceptionOccurrenceIdInstanceId, options ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions) (result ListEventExceptionOccurrenceInstanceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEventExceptionOccurrenceInstanceAttachmentsCustomPager{},
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

// ListEventExceptionOccurrenceInstanceAttachmentsComplete retrieves all the results into a single object
func (c EventExceptionOccurrenceInstanceAttachmentClient) ListEventExceptionOccurrenceInstanceAttachmentsComplete(ctx context.Context, id beta.MeEventIdExceptionOccurrenceIdInstanceId, options ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions) (ListEventExceptionOccurrenceInstanceAttachmentsCompleteResult, error) {
	return c.ListEventExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentOperationPredicate{})
}

// ListEventExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventExceptionOccurrenceInstanceAttachmentClient) ListEventExceptionOccurrenceInstanceAttachmentsCompleteMatchingPredicate(ctx context.Context, id beta.MeEventIdExceptionOccurrenceIdInstanceId, options ListEventExceptionOccurrenceInstanceAttachmentsOperationOptions, predicate AttachmentOperationPredicate) (result ListEventExceptionOccurrenceInstanceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListEventExceptionOccurrenceInstanceAttachments(ctx, id, options)
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

	result = ListEventExceptionOccurrenceInstanceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
