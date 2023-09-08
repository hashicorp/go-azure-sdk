package groupteamchannelsharedwithteam

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

type ListGroupByIdTeamChannelByIdSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SharedWithChannelTeamInfoCollectionResponse
}

type ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteResult struct {
	Items []models.SharedWithChannelTeamInfoCollectionResponse
}

// ListGroupByIdTeamChannelByIdSharedWithTeams ...
func (c GroupTeamChannelSharedWithTeamClient) ListGroupByIdTeamChannelByIdSharedWithTeams(ctx context.Context, id GroupTeamChannelId) (result ListGroupByIdTeamChannelByIdSharedWithTeamsOperationResponse, err error) {
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

// ListGroupByIdTeamChannelByIdSharedWithTeamsComplete retrieves all the results into a single object
func (c GroupTeamChannelSharedWithTeamClient) ListGroupByIdTeamChannelByIdSharedWithTeamsComplete(ctx context.Context, id GroupTeamChannelId) (ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteResult, error) {
	return c.ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteMatchingPredicate(ctx, id, models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamChannelSharedWithTeamClient) ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id GroupTeamChannelId, predicate models.SharedWithChannelTeamInfoCollectionResponseOperationPredicate) (result ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteResult, err error) {
	items := make([]models.SharedWithChannelTeamInfoCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamChannelByIdSharedWithTeams(ctx, id)
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

	result = ListGroupByIdTeamChannelByIdSharedWithTeamsCompleteResult{
		Items: items,
	}
	return
}
