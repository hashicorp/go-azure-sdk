package mejoinedteamtagmember

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

type ListMeJoinedTeamByIdTagByIdMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamworkTagMemberCollectionResponse
}

type ListMeJoinedTeamByIdTagByIdMembersCompleteResult struct {
	Items []models.TeamworkTagMemberCollectionResponse
}

// ListMeJoinedTeamByIdTagByIdMembers ...
func (c MeJoinedTeamTagMemberClient) ListMeJoinedTeamByIdTagByIdMembers(ctx context.Context, id MeJoinedTeamTagId) (result ListMeJoinedTeamByIdTagByIdMembersOperationResponse, err error) {
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

// ListMeJoinedTeamByIdTagByIdMembersComplete retrieves all the results into a single object
func (c MeJoinedTeamTagMemberClient) ListMeJoinedTeamByIdTagByIdMembersComplete(ctx context.Context, id MeJoinedTeamTagId) (ListMeJoinedTeamByIdTagByIdMembersCompleteResult, error) {
	return c.ListMeJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate(ctx, id, models.TeamworkTagMemberCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamTagMemberClient) ListMeJoinedTeamByIdTagByIdMembersCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamTagId, predicate models.TeamworkTagMemberCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdTagByIdMembersCompleteResult, err error) {
	items := make([]models.TeamworkTagMemberCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdTagByIdMembers(ctx, id)
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

	result = ListMeJoinedTeamByIdTagByIdMembersCompleteResult{
		Items: items,
	}
	return
}
