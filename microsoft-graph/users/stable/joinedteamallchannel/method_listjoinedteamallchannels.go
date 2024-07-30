package joinedteamallchannel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListJoinedTeamAllChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Channel
}

type ListJoinedTeamAllChannelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Channel
}

type ListJoinedTeamAllChannelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamAllChannelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamAllChannels ...
func (c JoinedTeamAllChannelClient) ListJoinedTeamAllChannels(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamAllChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamAllChannelsCustomPager{},
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
		Values *[]stable.Channel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamAllChannelsComplete retrieves all the results into a single object
func (c JoinedTeamAllChannelClient) ListJoinedTeamAllChannelsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamAllChannelsCompleteResult, error) {
	return c.ListJoinedTeamAllChannelsCompleteMatchingPredicate(ctx, id, ChannelOperationPredicate{})
}

// ListJoinedTeamAllChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamAllChannelClient) ListJoinedTeamAllChannelsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate ChannelOperationPredicate) (result ListJoinedTeamAllChannelsCompleteResult, err error) {
	items := make([]stable.Channel, 0)

	resp, err := c.ListJoinedTeamAllChannels(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListJoinedTeamAllChannelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
