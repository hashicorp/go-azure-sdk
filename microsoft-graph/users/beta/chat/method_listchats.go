package chat

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

type ListChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Chat
}

type ListChatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Chat
}

type ListChatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChats ...
func (c ChatClient) ListChats(ctx context.Context, id UserId) (result ListChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChatsCustomPager{},
		Path:       fmt.Sprintf("%s/chats", id.ID()),
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
		Values *[]beta.Chat `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatsComplete retrieves all the results into a single object
func (c ChatClient) ListChatsComplete(ctx context.Context, id UserId) (ListChatsCompleteResult, error) {
	return c.ListChatsCompleteMatchingPredicate(ctx, id, ChatOperationPredicate{})
}

// ListChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatClient) ListChatsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ChatOperationPredicate) (result ListChatsCompleteResult, err error) {
	items := make([]beta.Chat, 0)

	resp, err := c.ListChats(ctx, id)
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

	result = ListChatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
