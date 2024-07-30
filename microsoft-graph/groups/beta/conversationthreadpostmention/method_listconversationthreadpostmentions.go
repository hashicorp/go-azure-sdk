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

type ListConversationThreadPostMentionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostMentionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPostMentions ...
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentions(ctx context.Context, id GroupIdConversationIdThreadIdPostId) (result ListConversationThreadPostMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConversationThreadPostMentionsCustomPager{},
		Path:       fmt.Sprintf("%s/mentions", id.ID()),
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
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentionsComplete(ctx context.Context, id GroupIdConversationIdThreadIdPostId) (ListConversationThreadPostMentionsCompleteResult, error) {
	return c.ListConversationThreadPostMentionsCompleteMatchingPredicate(ctx, id, MentionOperationPredicate{})
}

// ListConversationThreadPostMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostMentionClient) ListConversationThreadPostMentionsCompleteMatchingPredicate(ctx context.Context, id GroupIdConversationIdThreadIdPostId, predicate MentionOperationPredicate) (result ListConversationThreadPostMentionsCompleteResult, err error) {
	items := make([]beta.Mention, 0)

	resp, err := c.ListConversationThreadPostMentions(ctx, id)
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
