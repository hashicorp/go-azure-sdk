package conversationthreadpostmention

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

type ListConversationThreadPostMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Mention
}

type ListConversationThreadPostMentionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Mention
}

type ListConversationThreadPostMentionsOperationOptions struct {
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

func DefaultListConversationThreadPostMentionsOperationOptions() ListConversationThreadPostMentionsOperationOptions {
	return ListConversationThreadPostMentionsOperationOptions{}
}

func (o ListConversationThreadPostMentionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConversationThreadPostMentionsOperationOptions) ToOData() *odata.Query {
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

func (o ListConversationThreadPostMentionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConversationThreadPostMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPostMentions - Get mentions from groups
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentions(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostMentionsOperationOptions) (result ListConversationThreadPostMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConversationThreadPostMentionsCustomPager{},
		Path:          fmt.Sprintf("%s/mentions", id.ID()),
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
		Values *[]beta.Mention `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConversationThreadPostMentionsComplete retrieves all the results into a single object
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentionsComplete(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostMentionsOperationOptions) (ListConversationThreadPostMentionsCompleteResult, error) {
	return c.ListConversationThreadPostMentionsCompleteMatchingPredicate(ctx, id, options, MentionOperationPredicate{})
}

// ListConversationThreadPostMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options ListConversationThreadPostMentionsOperationOptions, predicate MentionOperationPredicate) (result ListConversationThreadPostMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListConversationThreadPostMentions(ctx, id, options)
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

	result = ListConversationThreadPostMentionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
