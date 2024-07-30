package joinedteamincomingchannel

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

type ListJoinedTeamIncomingChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Channel
}

type ListJoinedTeamIncomingChannelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Channel
}

type ListJoinedTeamIncomingChannelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamIncomingChannelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamIncomingChannels ...
func (c JoinedTeamIncomingChannelClient) ListJoinedTeamIncomingChannels(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamIncomingChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamIncomingChannelsCustomPager{},
		Path:       fmt.Sprintf("%s/incomingChannels", id.ID()),
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

// ListJoinedTeamIncomingChannelsComplete retrieves all the results into a single object
func (c JoinedTeamIncomingChannelClient) ListJoinedTeamIncomingChannelsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamIncomingChannelsCompleteResult, error) {
	return c.ListJoinedTeamIncomingChannelsCompleteMatchingPredicate(ctx, id, ChannelOperationPredicate{})
}

// ListJoinedTeamIncomingChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamIncomingChannelClient) ListJoinedTeamIncomingChannelsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate ChannelOperationPredicate) (result ListJoinedTeamIncomingChannelsCompleteResult, err error) {
	items := make([]stable.Channel, 0)

	resp, err := c.ListJoinedTeamIncomingChannels(ctx, id)
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

	result = ListJoinedTeamIncomingChannelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
