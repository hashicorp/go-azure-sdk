package teamchannelmessagereply

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

type ListTeamChannelMessageRepliesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChatMessage
}

type ListTeamChannelMessageRepliesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChatMessage
}

type ListTeamChannelMessageRepliesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelMessageRepliesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelMessageReplies ...
func (c TeamChannelMessageReplyClient) ListTeamChannelMessageReplies(ctx context.Context, id GroupIdTeamChannelIdMessageId) (result ListTeamChannelMessageRepliesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamChannelMessageRepliesCustomPager{},
		Path:       fmt.Sprintf("%s/replies", id.ID()),
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
		Values *[]beta.ChatMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelMessageRepliesComplete retrieves all the results into a single object
func (c TeamChannelMessageReplyClient) ListTeamChannelMessageRepliesComplete(ctx context.Context, id GroupIdTeamChannelIdMessageId) (ListTeamChannelMessageRepliesCompleteResult, error) {
	return c.ListTeamChannelMessageRepliesCompleteMatchingPredicate(ctx, id, ChatMessageOperationPredicate{})
}

// ListTeamChannelMessageRepliesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMessageReplyClient) ListTeamChannelMessageRepliesCompleteMatchingPredicate(ctx context.Context, id GroupIdTeamChannelIdMessageId, predicate ChatMessageOperationPredicate) (result ListTeamChannelMessageRepliesCompleteResult, err error) {
	items := make([]beta.ChatMessage, 0)

	resp, err := c.ListTeamChannelMessageReplies(ctx, id)
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

	result = ListTeamChannelMessageRepliesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
