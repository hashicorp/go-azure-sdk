package userjoinedteamchannel

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

type ListUserByIdJoinedTeamByIdChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChannelCollectionResponse
}

type ListUserByIdJoinedTeamByIdChannelsCompleteResult struct {
	Items []models.ChannelCollectionResponse
}

// ListUserByIdJoinedTeamByIdChannels ...
func (c UserJoinedTeamChannelClient) ListUserByIdJoinedTeamByIdChannels(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/channels", id.ID()),
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

// ListUserByIdJoinedTeamByIdChannelsComplete retrieves all the results into a single object
func (c UserJoinedTeamChannelClient) ListUserByIdJoinedTeamByIdChannelsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdChannelsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdChannelsCompleteMatchingPredicate(ctx, id, models.ChannelCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamChannelClient) ListUserByIdJoinedTeamByIdChannelsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.ChannelCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdChannelsCompleteResult, err error) {
	items := make([]models.ChannelCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdChannels(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdChannelsCompleteResult{
		Items: items,
	}
	return
}
