package teamallchannel

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

type ListTeamAllChannelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Channel
}

type ListTeamAllChannelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Channel
}

type ListTeamAllChannelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamAllChannelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamAllChannels ...
func (c TeamAllChannelClient) ListTeamAllChannels(ctx context.Context, id GroupId) (result ListTeamAllChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamAllChannelsCustomPager{},
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
		Values *[]beta.Channel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamAllChannelsComplete retrieves all the results into a single object
func (c TeamAllChannelClient) ListTeamAllChannelsComplete(ctx context.Context, id GroupId) (ListTeamAllChannelsCompleteResult, error) {
	return c.ListTeamAllChannelsCompleteMatchingPredicate(ctx, id, ChannelOperationPredicate{})
}

// ListTeamAllChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamAllChannelClient) ListTeamAllChannelsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate ChannelOperationPredicate) (result ListTeamAllChannelsCompleteResult, err error) {
	items := make([]beta.Channel, 0)

	resp, err := c.ListTeamAllChannels(ctx, id)
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

	result = ListTeamAllChannelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
