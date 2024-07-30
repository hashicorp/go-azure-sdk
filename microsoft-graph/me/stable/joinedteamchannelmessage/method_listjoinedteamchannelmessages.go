package joinedteamchannelmessage

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

type ListJoinedTeamChannelMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessage
}

type ListJoinedTeamChannelMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessage
}

type ListJoinedTeamChannelMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamChannelMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamChannelMessages ...
func (c JoinedTeamChannelMessageClient) ListJoinedTeamChannelMessages(ctx context.Context, id MeJoinedTeamIdChannelId) (result ListJoinedTeamChannelMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamChannelMessagesCustomPager{},
		Path:       fmt.Sprintf("%s/messages", id.ID()),
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
		Values *[]stable.ChatMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamChannelMessagesComplete retrieves all the results into a single object
func (c JoinedTeamChannelMessageClient) ListJoinedTeamChannelMessagesComplete(ctx context.Context, id MeJoinedTeamIdChannelId) (ListJoinedTeamChannelMessagesCompleteResult, error) {
	return c.ListJoinedTeamChannelMessagesCompleteMatchingPredicate(ctx, id, ChatMessageOperationPredicate{})
}

// ListJoinedTeamChannelMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelMessageClient) ListJoinedTeamChannelMessagesCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamIdChannelId, predicate ChatMessageOperationPredicate) (result ListJoinedTeamChannelMessagesCompleteResult, err error) {
	items := make([]stable.ChatMessage, 0)

	resp, err := c.ListJoinedTeamChannelMessages(ctx, id)
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

	result = ListJoinedTeamChannelMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
