package mejoinedteamprimarychannelsharedwithteamallowedmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembers ...
func (c MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembers(ctx context.Context, id MeJoinedTeamPrimaryChannelSharedWithTeamId) (result ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.ConversationMemberCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersComplete retrieves all the results into a single object
func (c MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersComplete(ctx context.Context, id MeJoinedTeamPrimaryChannelSharedWithTeamId) (ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult, error) {
	return c.ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamPrimaryChannelSharedWithTeamId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembers(ctx, id)
	if err != nil {
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

	result = ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult{
		Items: items,
	}
	return
}
