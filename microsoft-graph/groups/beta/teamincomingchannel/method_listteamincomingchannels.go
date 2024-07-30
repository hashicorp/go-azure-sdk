package teamincomingchannel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTeamIncomingChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Channel
}

type ListTeamIncomingChannelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Channel
}

type ListTeamIncomingChannelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamIncomingChannelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamIncomingChannels ...
func (c TeamIncomingChannelClient) ListTeamIncomingChannels(ctx context.Context, id GroupId) (result ListTeamIncomingChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamIncomingChannelsCustomPager{},
		Path:       fmt.Sprintf("%s/team/incomingChannels", id.ID()),
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
		Values *[]beta.Channel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamIncomingChannelsComplete retrieves all the results into a single object
func (c TeamIncomingChannelClient) ListTeamIncomingChannelsComplete(ctx context.Context, id GroupId) (ListTeamIncomingChannelsCompleteResult, error) {
	return c.ListTeamIncomingChannelsCompleteMatchingPredicate(ctx, id, ChannelOperationPredicate{})
}

// ListTeamIncomingChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamIncomingChannelClient) ListTeamIncomingChannelsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate ChannelOperationPredicate) (result ListTeamIncomingChannelsCompleteResult, err error) {
	items := make([]beta.Channel, 0)

	resp, err := c.ListTeamIncomingChannels(ctx, id)
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

	result = ListTeamIncomingChannelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
