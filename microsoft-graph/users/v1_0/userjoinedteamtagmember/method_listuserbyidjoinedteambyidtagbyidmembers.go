package userjoinedteamtagmember

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

type ListUserByIdJoinedTeamByIdTagByIdMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamworkTagMemberCollectionResponse
}

type ListUserByIdJoinedTeamByIdTagByIdMembersCompleteResult struct {
	Items []models.TeamworkTagMemberCollectionResponse
}

// ListUserByIdJoinedTeamByIdTagByIdMembers ...
func (c UserJoinedTeamTagMemberClient) ListUserByIdJoinedTeamByIdTagByIdMembers(ctx context.Context, id UserJoinedTeamTagId) (result ListUserByIdJoinedTeamByIdTagByIdMembersOperationResponse, err error) {
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
		Values *[]models.TeamworkTagMemberCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdTagByIdMembersComplete retrieves all the results into a single object
func (c UserJoinedTeamTagMemberClient) ListUserByIdJoinedTeamByIdTagByIdMembersComplete(ctx context.Context, id UserJoinedTeamTagId) (ListUserByIdJoinedTeamByIdTagByIdMembersCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate(ctx, id, models.TeamworkTagMemberCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamTagMemberClient) ListUserByIdJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamTagId, predicate models.TeamworkTagMemberCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdTagByIdMembersCompleteResult, err error) {
	items := make([]models.TeamworkTagMemberCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdTagByIdMembers(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdTagByIdMembersCompleteResult{
		Items: items,
	}
	return
}
