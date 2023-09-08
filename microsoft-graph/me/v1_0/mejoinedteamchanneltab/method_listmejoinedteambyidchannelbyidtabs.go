package mejoinedteamchanneltab

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

type ListMeJoinedTeamByIdChannelByIdTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsTabCollectionResponse
}

type ListMeJoinedTeamByIdChannelByIdTabsCompleteResult struct {
	Items []models.TeamsTabCollectionResponse
}

// ListMeJoinedTeamByIdChannelByIdTabs ...
func (c MeJoinedTeamChannelTabClient) ListMeJoinedTeamByIdChannelByIdTabs(ctx context.Context, id MeJoinedTeamChannelId) (result ListMeJoinedTeamByIdChannelByIdTabsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdChannelByIdTabsComplete retrieves all the results into a single object
func (c MeJoinedTeamChannelTabClient) ListMeJoinedTeamByIdChannelByIdTabsComplete(ctx context.Context, id MeJoinedTeamChannelId) (ListMeJoinedTeamByIdChannelByIdTabsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate(ctx, id, models.TeamsTabCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamChannelTabClient) ListMeJoinedTeamByIdChannelByIdTabsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamChannelId, predicate models.TeamsTabCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdChannelByIdTabsCompleteResult, err error) {
	items := make([]models.TeamsTabCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdChannelByIdTabs(ctx, id)
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

	result = ListMeJoinedTeamByIdChannelByIdTabsCompleteResult{
		Items: items,
	}
	return
}
