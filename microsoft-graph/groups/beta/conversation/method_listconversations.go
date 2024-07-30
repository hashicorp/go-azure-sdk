package conversation

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

type ListConversationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Conversation
}

type ListConversationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Conversation
}

type ListConversationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversations ...
func (c ConversationClient) ListConversations(ctx context.Context, id GroupId) (result ListConversationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConversationsCustomPager{},
		Path:       fmt.Sprintf("%s/conversations", id.ID()),
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
		Values *[]beta.Conversation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConversationsComplete retrieves all the results into a single object
func (c ConversationClient) ListConversationsComplete(ctx context.Context, id GroupId) (ListConversationsCompleteResult, error) {
	return c.ListConversationsCompleteMatchingPredicate(ctx, id, ConversationOperationPredicate{})
}

// ListConversationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationClient) ListConversationsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate ConversationOperationPredicate) (result ListConversationsCompleteResult, err error) {
	items := make([]beta.Conversation, 0)

	resp, err := c.ListConversations(ctx, id)
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

	result = ListConversationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
