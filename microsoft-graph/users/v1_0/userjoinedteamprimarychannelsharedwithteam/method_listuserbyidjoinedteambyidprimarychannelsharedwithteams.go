package userjoinedteamprimarychannelsharedwithteam

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

type ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedWithChannelTeamInfoCollectionResponse
}

type ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult struct {
	Items []models.SharedWithChannelTeamInfoCollectionResponse
}

// ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeams ...
func (c UserJoinedTeamPrimaryChannelSharedWithTeamClient) ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeams(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/primaryChannel/sharedWithTeams", id.ID()),
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

// ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c UserJoinedTeamPrimaryChannelSharedWithTeamClient) ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamPrimaryChannelSharedWithTeamClient) ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]models.SharedWithChannelTeamInfoCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeams(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult{
		Items: items,
	}
	return
}
