package conversationthreadpost

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

type ListConversationThreadPostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Post
}

type ListConversationThreadPostsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Post
}

type ListConversationThreadPostsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadPostsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreadPosts ...
func (c ConversationThreadPostClient) ListConversationThreadPosts(ctx context.Context, id GroupIdConversationIdThreadId) (result ListConversationThreadPostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConversationThreadPostsCustomPager{},
		Path:       fmt.Sprintf("%s/posts", id.ID()),
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
		Values *[]stable.Post `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConversationThreadPostsComplete retrieves all the results into a single object
func (c ConversationThreadPostClient) ListConversationThreadPostsComplete(ctx context.Context, id GroupIdConversationIdThreadId) (ListConversationThreadPostsCompleteResult, error) {
	return c.ListConversationThreadPostsCompleteMatchingPredicate(ctx, id, PostOperationPredicate{})
}

// ListConversationThreadPostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadPostClient) ListConversationThreadPostsCompleteMatchingPredicate(ctx context.Context, id GroupIdConversationIdThreadId, predicate PostOperationPredicate) (result ListConversationThreadPostsCompleteResult, err error) {
	items := make([]stable.Post, 0)

	resp, err := c.ListConversationThreadPosts(ctx, id)
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

	result = ListConversationThreadPostsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
