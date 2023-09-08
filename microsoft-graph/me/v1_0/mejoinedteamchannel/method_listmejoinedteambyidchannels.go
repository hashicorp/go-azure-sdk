package mejoinedteamchannel

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

type ListMeJoinedTeamByIdChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChannelCollectionResponse
}

type ListMeJoinedTeamByIdChannelsCompleteResult struct {
	Items []models.ChannelCollectionResponse
}

// ListMeJoinedTeamByIdChannels ...
func (c MeJoinedTeamChannelClient) ListMeJoinedTeamByIdChannels(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdChannelsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdChannelsComplete retrieves all the results into a single object
func (c MeJoinedTeamChannelClient) ListMeJoinedTeamByIdChannelsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdChannelsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdChannelsCompleteMatchingPredicate(ctx, id, models.ChannelCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamChannelClient) ListMeJoinedTeamByIdChannelsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.ChannelCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdChannelsCompleteResult, err error) {
	items := make([]models.ChannelCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdChannels(ctx, id)
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

	result = ListMeJoinedTeamByIdChannelsCompleteResult{
		Items: items,
	}
	return
}
