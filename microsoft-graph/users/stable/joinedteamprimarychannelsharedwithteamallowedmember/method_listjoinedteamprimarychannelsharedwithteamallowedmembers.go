package joinedteamprimarychannelsharedwithteamallowedmember

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

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers ...
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers(ctx context.Context, id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) (result ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager{},
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

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersComplete(ctx context.Context, id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId) (ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx, id, ConversationMemberOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, predicate ConversationMemberOperationPredicate) (result ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers(ctx, id)
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

	result = ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
