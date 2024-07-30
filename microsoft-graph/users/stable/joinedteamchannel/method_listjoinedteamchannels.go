package joinedteamchannel

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

type ListJoinedTeamChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Channel
}

type ListJoinedTeamChannelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Channel
}

type ListJoinedTeamChannelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamChannelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamChannels ...
func (c JoinedTeamChannelClient) ListJoinedTeamChannels(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamChannelsCustomPager{},
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
		Values *[]stable.Channel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamChannelsComplete retrieves all the results into a single object
func (c JoinedTeamChannelClient) ListJoinedTeamChannelsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamChannelsCompleteResult, error) {
	return c.ListJoinedTeamChannelsCompleteMatchingPredicate(ctx, id, ChannelOperationPredicate{})
}

// ListJoinedTeamChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelClient) ListJoinedTeamChannelsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate ChannelOperationPredicate) (result ListJoinedTeamChannelsCompleteResult, err error) {
	items := make([]stable.Channel, 0)

	resp, err := c.ListJoinedTeamChannels(ctx, id)
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

	result = ListJoinedTeamChannelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
