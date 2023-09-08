package userjoinedteamallchannel

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

type ListUserByIdJoinedTeamByIdAllChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChannelCollectionResponse
}

type ListUserByIdJoinedTeamByIdAllChannelsCompleteResult struct {
	Items []models.ChannelCollectionResponse
}

// ListUserByIdJoinedTeamByIdAllChannels ...
func (c UserJoinedTeamAllChannelClient) ListUserByIdJoinedTeamByIdAllChannels(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdAllChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/allChannels", id.ID()),
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
		Values *[]models.ChannelCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdAllChannelsComplete retrieves all the results into a single object
func (c UserJoinedTeamAllChannelClient) ListUserByIdJoinedTeamByIdAllChannelsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdAllChannelsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdAllChannelsCompleteMatchingPredicate(ctx, id, models.ChannelCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdAllChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamAllChannelClient) ListUserByIdJoinedTeamByIdAllChannelsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.ChannelCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdAllChannelsCompleteResult, err error) {
	items := make([]models.ChannelCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdAllChannels(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdAllChannelsCompleteResult{
		Items: items,
	}
	return
}
