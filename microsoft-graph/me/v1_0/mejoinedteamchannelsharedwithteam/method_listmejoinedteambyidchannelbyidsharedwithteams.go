package mejoinedteamchannelsharedwithteam

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

type ListMeJoinedTeamByIdChannelByIdSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedWithChannelTeamInfoCollectionResponse
}

type ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteResult struct {
	Items []models.SharedWithChannelTeamInfoCollectionResponse
}

// ListMeJoinedTeamByIdChannelByIdSharedWithTeams ...
func (c MeJoinedTeamChannelSharedWithTeamClient) ListMeJoinedTeamByIdChannelByIdSharedWithTeams(ctx context.Context, id MeJoinedTeamChannelId) (result ListMeJoinedTeamByIdChannelByIdSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sharedWithTeams", id.ID()),
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
		Values *[]models.SharedWithChannelTeamInfoCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeJoinedTeamByIdChannelByIdSharedWithTeamsComplete retrieves all the results into a single object
func (c MeJoinedTeamChannelSharedWithTeamClient) ListMeJoinedTeamByIdChannelByIdSharedWithTeamsComplete(ctx context.Context, id MeJoinedTeamChannelId) (ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteMatchingPredicate(ctx, id, models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamChannelSharedWithTeamClient) ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamChannelId, predicate models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteResult, err error) {
	items := make([]models.SharedWithChannelTeamInfoCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdChannelByIdSharedWithTeams(ctx, id)
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

	result = ListMeJoinedTeamByIdChannelByIdSharedWithTeamsCompleteResult{
		Items: items,
	}
	return
}
