package groupteamprimarychannelsharedwithteamallowedmember

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

type ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembers ...
func (c GroupTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembers(ctx context.Context, id GroupTeamPrimaryChannelSharedWithTeamId) (result ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersOperationResponse, err error) {
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

// ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersComplete retrieves all the results into a single object
func (c GroupTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersComplete(ctx context.Context, id GroupTeamPrimaryChannelSharedWithTeamId) (ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult, error) {
	return c.ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteMatchingPredicate(ctx context.Context, id GroupTeamPrimaryChannelSharedWithTeamId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembers(ctx, id)
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

	result = ListGroupByIdTeamPrimaryChannelSharedWithTeamByIdAllowedMembersCompleteResult{
		Items: items,
	}
	return
}
