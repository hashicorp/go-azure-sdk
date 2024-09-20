package conversationthreadpostattachment

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

type ListConversationThreadPostAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListConversationThreadPostAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListConversationThreadPostAttachmentsOperationOptions struct {
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

func DefaultListConversationThreadPostAttachmentsOperationOptions() ListConversationThreadPostAttachmentsOperationOptions {
	return ListConversationThreadPostAttachmentsOperationOptions{}
}

func (o ListConversationThreadPostAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConversationThreadPostAttachmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListConversationThreadPostAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConversationThreadPostAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPostAttachments - Get attachments from groups. The collection of fileAttachment,
// itemAttachment, and referenceAttachment attachments for the post. Read-only. Nullable. Supports $expand.
func (c ConversationThreadPostAttachmentClient) ListConversationThreadPostAttachments(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostAttachmentsOperationOptions) (result ListConversationThreadPostAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConversationThreadPostAttachmentsCustomPager{},
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

// ListConversationThreadPostAttachmentsComplete retrieves all the results into a single object
func (c ConversationThreadPostAttachmentClient) ListConversationThreadPostAttachmentsComplete(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostAttachmentsOperationOptions) (ListConversationThreadPostAttachmentsCompleteResult, error) {
	return c.ListConversationThreadPostAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentOperationPredicate{})
}

// ListConversationThreadPostAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostAttachmentClient) ListConversationThreadPostAttachmentsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostAttachmentsOperationOptions, predicate AttachmentOperationPredicate) (result ListConversationThreadPostAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListConversationThreadPostAttachments(ctx, id, options)
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

	result = ListConversationThreadPostAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
