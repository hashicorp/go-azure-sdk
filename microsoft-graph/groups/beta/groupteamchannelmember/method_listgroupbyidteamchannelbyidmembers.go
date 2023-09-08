package groupteamchannelmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdTeamChannelByIdMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListGroupByIdTeamChannelByIdMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListGroupByIdTeamChannelByIdMembers ...
func (c GroupTeamChannelMemberClient) ListGroupByIdTeamChannelByIdMembers(ctx context.Context, id GroupTeamChannelId) (result ListGroupByIdTeamChannelByIdMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.ConversationMemberCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamChannelByIdMembersComplete retrieves all the results into a single object
func (c GroupTeamChannelMemberClient) ListGroupByIdTeamChannelByIdMembersComplete(ctx context.Context, id GroupTeamChannelId) (ListGroupByIdTeamChannelByIdMembersCompleteResult, error) {
	return c.ListGroupByIdTeamChannelByIdMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamChannelByIdMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamChannelMemberClient) ListGroupByIdTeamChannelByIdMembersCompleteMatchingPredicate(ctx context.Context, id GroupTeamChannelId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListGroupByIdTeamChannelByIdMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamChannelByIdMembers(ctx, id)
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

	result = ListGroupByIdTeamChannelByIdMembersCompleteResult{
		Items: items,
	}
	return
}
