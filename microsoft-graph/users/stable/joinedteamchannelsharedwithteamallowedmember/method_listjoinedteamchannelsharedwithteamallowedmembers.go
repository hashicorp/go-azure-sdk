package joinedteamchannelsharedwithteamallowedmember

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

type ListJoinedTeamChannelSharedWithTeamAllowedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListJoinedTeamChannelSharedWithTeamAllowedMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamChannelSharedWithTeamAllowedMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamChannelSharedWithTeamAllowedMembers ...
func (c JoinedTeamChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamChannelSharedWithTeamAllowedMembers(ctx context.Context, id UserIdJoinedTeamIdChannelIdSharedWithTeamId) (result ListJoinedTeamChannelSharedWithTeamAllowedMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamChannelSharedWithTeamAllowedMembersCustomPager{},
		Path:       fmt.Sprintf("%s/allowedMembers", id.ID()),
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
		Values *[]stable.ConversationMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamChannelSharedWithTeamAllowedMembersComplete retrieves all the results into a single object
func (c JoinedTeamChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamChannelSharedWithTeamAllowedMembersComplete(ctx context.Context, id UserIdJoinedTeamIdChannelIdSharedWithTeamId) (ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteResult, error) {
	return c.ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx, id, ConversationMemberOperationPredicate{})
}

// ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamIdChannelIdSharedWithTeamId, predicate ConversationMemberOperationPredicate) (result ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListJoinedTeamChannelSharedWithTeamAllowedMembers(ctx, id)
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

	result = ListJoinedTeamChannelSharedWithTeamAllowedMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
