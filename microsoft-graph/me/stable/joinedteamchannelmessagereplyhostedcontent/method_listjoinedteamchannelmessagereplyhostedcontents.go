package joinedteamchannelmessagereplyhostedcontent

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

type ListJoinedTeamChannelMessageReplyHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessageHostedContent
}

type ListJoinedTeamChannelMessageReplyHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessageHostedContent
}

type ListJoinedTeamChannelMessageReplyHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamChannelMessageReplyHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamChannelMessageReplyHostedContents ...
func (c JoinedTeamChannelMessageReplyHostedContentClient) ListJoinedTeamChannelMessageReplyHostedContents(ctx context.Context, id MeJoinedTeamIdChannelIdMessageIdReplyId) (result ListJoinedTeamChannelMessageReplyHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamChannelMessageReplyHostedContentsCustomPager{},
		Path:       fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]stable.ChatMessageHostedContent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamChannelMessageReplyHostedContentsComplete retrieves all the results into a single object
func (c JoinedTeamChannelMessageReplyHostedContentClient) ListJoinedTeamChannelMessageReplyHostedContentsComplete(ctx context.Context, id MeJoinedTeamIdChannelIdMessageIdReplyId) (ListJoinedTeamChannelMessageReplyHostedContentsCompleteResult, error) {
	return c.ListJoinedTeamChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx, id, ChatMessageHostedContentOperationPredicate{})
}

// ListJoinedTeamChannelMessageReplyHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelMessageReplyHostedContentClient) ListJoinedTeamChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamIdChannelIdMessageIdReplyId, predicate ChatMessageHostedContentOperationPredicate) (result ListJoinedTeamChannelMessageReplyHostedContentsCompleteResult, err error) {
	items := make([]stable.ChatMessageHostedContent, 0)

	resp, err := c.ListJoinedTeamChannelMessageReplyHostedContents(ctx, id)
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

	result = ListJoinedTeamChannelMessageReplyHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
