package chatpinnedmessage

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

type ListChatPinnedMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PinnedChatMessageInfo
}

type ListChatPinnedMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PinnedChatMessageInfo
}

type ListChatPinnedMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatPinnedMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatPinnedMessages ...
func (c ChatPinnedMessageClient) ListChatPinnedMessages(ctx context.Context, id MeChatId) (result ListChatPinnedMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChatPinnedMessagesCustomPager{},
		Path:       fmt.Sprintf("%s/pinnedMessages", id.ID()),
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
		Values *[]stable.PinnedChatMessageInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatPinnedMessagesComplete retrieves all the results into a single object
func (c ChatPinnedMessageClient) ListChatPinnedMessagesComplete(ctx context.Context, id MeChatId) (ListChatPinnedMessagesCompleteResult, error) {
	return c.ListChatPinnedMessagesCompleteMatchingPredicate(ctx, id, PinnedChatMessageInfoOperationPredicate{})
}

// ListChatPinnedMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatPinnedMessageClient) ListChatPinnedMessagesCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate PinnedChatMessageInfoOperationPredicate) (result ListChatPinnedMessagesCompleteResult, err error) {
	items := make([]stable.PinnedChatMessageInfo, 0)

	resp, err := c.ListChatPinnedMessages(ctx, id)
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

	result = ListChatPinnedMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
