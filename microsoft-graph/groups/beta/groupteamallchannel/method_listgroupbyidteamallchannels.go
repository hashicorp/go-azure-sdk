package groupteamallchannel

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

type ListGroupByIdTeamAllChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChannelCollectionResponse
}

type ListGroupByIdTeamAllChannelsCompleteResult struct {
	Items []models.ChannelCollectionResponse
}

// ListGroupByIdTeamAllChannels ...
func (c GroupTeamAllChannelClient) ListGroupByIdTeamAllChannels(ctx context.Context, id GroupId) (result ListGroupByIdTeamAllChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/allChannels", id.ID()),
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

// ListGroupByIdTeamAllChannelsComplete retrieves all the results into a single object
func (c GroupTeamAllChannelClient) ListGroupByIdTeamAllChannelsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamAllChannelsCompleteResult, error) {
	return c.ListGroupByIdTeamAllChannelsCompleteMatchingPredicate(ctx, id, models.ChannelCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamAllChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamAllChannelClient) ListGroupByIdTeamAllChannelsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ChannelCollectionResponseOperationPredicate) (result ListGroupByIdTeamAllChannelsCompleteResult, err error) {
	items := make([]models.ChannelCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamAllChannels(ctx, id)
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

	result = ListGroupByIdTeamAllChannelsCompleteResult{
		Items: items,
	}
	return
}
