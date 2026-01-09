package teamincomingchannel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

type ListTeamIncomingChannelsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListTeamIncomingChannelsOperationOptions() ListTeamIncomingChannelsOperationOptions {
	return ListTeamIncomingChannelsOperationOptions{}
}

func (o ListTeamIncomingChannelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamIncomingChannelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTeamIncomingChannelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListTeamIncomingChannels - Get incomingChannels from groups. List of channels shared with the team.
func (c TeamIncomingChannelClient) ListTeamIncomingChannels(ctx context.Context, id beta.GroupId, options ListTeamIncomingChannelsOperationOptions) (result ListTeamIncomingChannelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamIncomingChannelsCustomPager{},
		Path:          fmt.Sprintf("%s/team/incomingChannels", id.ID()),
		RetryFunc:     options.RetryFunc,
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
func (c TeamIncomingChannelClient) ListTeamIncomingChannelsComplete(ctx context.Context, id beta.GroupId, options ListTeamIncomingChannelsOperationOptions) (ListTeamIncomingChannelsCompleteResult, error) {
	return c.ListTeamIncomingChannelsCompleteMatchingPredicate(ctx, id, options, ChannelOperationPredicate{})
}

// ListTeamIncomingChannelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamIncomingChannelClient) ListTeamIncomingChannelsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTeamIncomingChannelsOperationOptions, predicate ChannelOperationPredicate) (result ListTeamIncomingChannelsCompleteResult, err error) {
	items := make([]beta.Channel, 0)

	resp, err := c.ListTeamIncomingChannels(ctx, id, options)
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
