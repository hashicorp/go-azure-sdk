package conversationthreadpostinreplytomention

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

type ListConversationThreadPostInReplyToMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Mention
}

type ListConversationThreadPostInReplyToMentionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Mention
}

type ListConversationThreadPostInReplyToMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostInReplyToMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPostInReplyToMentions ...
func (c ConversationThreadPostInReplyToMentionClient) ListConversationThreadPostInReplyToMentions(ctx context.Context, id GroupIdConversationIdThreadIdPostId) (result ListConversationThreadPostInReplyToMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConversationThreadPostInReplyToMentionsCustomPager{},
		Path:       fmt.Sprintf("%s/inReplyTo/mentions", id.ID()),
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

// ListConversationThreadPostInReplyToMentionsComplete retrieves all the results into a single object
func (c ConversationThreadPostInReplyToMentionClient) ListConversationThreadPostInReplyToMentionsComplete(ctx context.Context, id GroupIdConversationIdThreadIdPostId) (ListConversationThreadPostInReplyToMentionsCompleteResult, error) {
	return c.ListConversationThreadPostInReplyToMentionsCompleteMatchingPredicate(ctx, id, MentionOperationPredicate{})
}

// ListConversationThreadPostInReplyToMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostInReplyToMentionClient) ListConversationThreadPostInReplyToMentionsCompleteMatchingPredicate(ctx context.Context, id GroupIdConversationIdThreadIdPostId, predicate MentionOperationPredicate) (result ListConversationThreadPostInReplyToMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListConversationThreadPostInReplyToMentions(ctx, id)
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

	result = ListConversationThreadPostInReplyToMentionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
