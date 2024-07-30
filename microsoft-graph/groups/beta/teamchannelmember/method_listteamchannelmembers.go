package teamchannelmember

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

type ListTeamChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConversationMember
}

type ListTeamChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConversationMember
}

type ListTeamChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelMembers ...
func (c TeamChannelMemberClient) ListTeamChannelMembers(ctx context.Context, id GroupIdTeamChannelId) (result ListTeamChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamChannelMembersCustomPager{},
		Path:       fmt.Sprintf("%s/members", id.ID()),
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
		Values *[]beta.ConversationMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelMembersComplete retrieves all the results into a single object
func (c TeamChannelMemberClient) ListTeamChannelMembersComplete(ctx context.Context, id GroupIdTeamChannelId) (ListTeamChannelMembersCompleteResult, error) {
	return c.ListTeamChannelMembersCompleteMatchingPredicate(ctx, id, ConversationMemberOperationPredicate{})
}

// ListTeamChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMemberClient) ListTeamChannelMembersCompleteMatchingPredicate(ctx context.Context, id GroupIdTeamChannelId, predicate ConversationMemberOperationPredicate) (result ListTeamChannelMembersCompleteResult, err error) {
	items := make([]beta.ConversationMember, 0)

	resp, err := c.ListTeamChannelMembers(ctx, id)
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

	result = ListTeamChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
