package userjoinedteamprimarychannelmember

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

type ListUserByIdJoinedTeamByIdPrimaryChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListUserByIdJoinedTeamByIdPrimaryChannelMembers ...
func (c UserJoinedTeamPrimaryChannelMemberClient) ListUserByIdJoinedTeamByIdPrimaryChannelMembers(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdPrimaryChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/primaryChannel/members", id.ID()),
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

// ListUserByIdJoinedTeamByIdPrimaryChannelMembersComplete retrieves all the results into a single object
func (c UserJoinedTeamPrimaryChannelMemberClient) ListUserByIdJoinedTeamByIdPrimaryChannelMembersComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamPrimaryChannelMemberClient) ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdPrimaryChannelMembers(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdPrimaryChannelMembersCompleteResult{
		Items: items,
	}
	return
}
