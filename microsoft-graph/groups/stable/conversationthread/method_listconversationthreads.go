package conversationthread

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

type ListConversationThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationThread
}

type ListConversationThreadsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationThread
}

type ListConversationThreadsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreads ...
func (c ConversationThreadClient) ListConversationThreads(ctx context.Context, id GroupIdConversationId) (result ListConversationThreadsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConversationThreadsCustomPager{},
		Path:       fmt.Sprintf("%s/threads", id.ID()),
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
		Values *[]stable.ConversationThread `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConversationThreadsComplete retrieves all the results into a single object
func (c ConversationThreadClient) ListConversationThreadsComplete(ctx context.Context, id GroupIdConversationId) (ListConversationThreadsCompleteResult, error) {
	return c.ListConversationThreadsCompleteMatchingPredicate(ctx, id, ConversationThreadOperationPredicate{})
}

// ListConversationThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadClient) ListConversationThreadsCompleteMatchingPredicate(ctx context.Context, id GroupIdConversationId, predicate ConversationThreadOperationPredicate) (result ListConversationThreadsCompleteResult, err error) {
	items := make([]stable.ConversationThread, 0)

	resp, err := c.ListConversationThreads(ctx, id)
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

	result = ListConversationThreadsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
