package mejoinedteamprimarychannelsharedwithteam

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

type ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedWithChannelTeamInfoCollectionResponse
}

type ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult struct {
	Items []models.SharedWithChannelTeamInfoCollectionResponse
}

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeams ...
func (c MeJoinedTeamPrimaryChannelSharedWithTeamClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeams(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c MeJoinedTeamPrimaryChannelSharedWithTeamClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamPrimaryChannelSharedWithTeamClient) ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]models.SharedWithChannelTeamInfoCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdPrimaryChannelSharedWithTeams(ctx, id)
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

	result = ListMeJoinedTeamByIdPrimaryChannelSharedWithTeamsCompleteResult{
		Items: items,
	}
	return
}
