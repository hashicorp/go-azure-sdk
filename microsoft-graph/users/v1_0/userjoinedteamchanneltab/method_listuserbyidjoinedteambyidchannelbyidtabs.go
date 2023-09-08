package userjoinedteamchanneltab

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

type ListUserByIdJoinedTeamByIdChannelByIdTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsTabCollectionResponse
}

type ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteResult struct {
	Items []models.TeamsTabCollectionResponse
}

// ListUserByIdJoinedTeamByIdChannelByIdTabs ...
func (c UserJoinedTeamChannelTabClient) ListUserByIdJoinedTeamByIdChannelByIdTabs(ctx context.Context, id UserJoinedTeamChannelId) (result ListUserByIdJoinedTeamByIdChannelByIdTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]models.TeamsTabCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdChannelByIdTabsComplete retrieves all the results into a single object
func (c UserJoinedTeamChannelTabClient) ListUserByIdJoinedTeamByIdChannelByIdTabsComplete(ctx context.Context, id UserJoinedTeamChannelId) (ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate(ctx, id, models.TeamsTabCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamChannelTabClient) ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamChannelId, predicate models.TeamsTabCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteResult, err error) {
	items := make([]models.TeamsTabCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdChannelByIdTabs(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdChannelByIdTabsCompleteResult{
		Items: items,
	}
	return
}
