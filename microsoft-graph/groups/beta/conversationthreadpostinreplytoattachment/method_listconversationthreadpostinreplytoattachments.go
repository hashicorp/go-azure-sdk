package conversationthreadpostinreplytoattachment

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

type ListConversationThreadPostInReplyToAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListConversationThreadPostInReplyToAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListConversationThreadPostInReplyToAttachmentsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListConversationThreadPostInReplyToAttachmentsOperationOptions() ListConversationThreadPostInReplyToAttachmentsOperationOptions {
	return ListConversationThreadPostInReplyToAttachmentsOperationOptions{}
}

func (o ListConversationThreadPostInReplyToAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConversationThreadPostInReplyToAttachmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListConversationThreadPostInReplyToAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConversationThreadPostInReplyToAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostInReplyToAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPostInReplyToAttachments - Get attachments from groups. The collection of fileAttachment,
// itemAttachment, and referenceAttachment attachments for the post. Read-only. Nullable. Supports $expand.
func (c ConversationThreadPostInReplyToAttachmentClient) ListConversationThreadPostInReplyToAttachments(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostInReplyToAttachmentsOperationOptions) (result ListConversationThreadPostInReplyToAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConversationThreadPostInReplyToAttachmentsCustomPager{},
		Path:          fmt.Sprintf("%s/inReplyTo/attachments", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListConversationThreadPostInReplyToAttachmentsComplete retrieves all the results into a single object
func (c ConversationThreadPostInReplyToAttachmentClient) ListConversationThreadPostInReplyToAttachmentsComplete(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostInReplyToAttachmentsOperationOptions) (ListConversationThreadPostInReplyToAttachmentsCompleteResult, error) {
	return c.ListConversationThreadPostInReplyToAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentOperationPredicate{})
}

// ListConversationThreadPostInReplyToAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostInReplyToAttachmentClient) ListConversationThreadPostInReplyToAttachmentsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostInReplyToAttachmentsOperationOptions, predicate AttachmentOperationPredicate) (result ListConversationThreadPostInReplyToAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListConversationThreadPostInReplyToAttachments(ctx, id, options)
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

	result = ListConversationThreadPostInReplyToAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
