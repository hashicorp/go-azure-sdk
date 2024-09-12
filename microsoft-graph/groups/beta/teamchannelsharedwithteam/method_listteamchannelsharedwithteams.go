package teamchannelsharedwithteam

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

type ListTeamChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedWithChannelTeamInfo
}

type ListTeamChannelSharedWithTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedWithChannelTeamInfo
}

type ListTeamChannelSharedWithTeamsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTeamChannelSharedWithTeamsOperationOptions() ListTeamChannelSharedWithTeamsOperationOptions {
	return ListTeamChannelSharedWithTeamsOperationOptions{}
}

func (o ListTeamChannelSharedWithTeamsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelSharedWithTeamsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelSharedWithTeamsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelSharedWithTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelSharedWithTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelSharedWithTeams - Get sharedWithTeams from groups. A collection of teams with which a channel is
// shared.
func (c TeamChannelSharedWithTeamClient) ListTeamChannelSharedWithTeams(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelSharedWithTeamsOperationOptions) (result ListTeamChannelSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelSharedWithTeamsCustomPager{},
		Path:          fmt.Sprintf("%s/sharedWithTeams", id.ID()),
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
		Values *[]beta.SharedWithChannelTeamInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c TeamChannelSharedWithTeamClient) ListTeamChannelSharedWithTeamsComplete(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelSharedWithTeamsOperationOptions) (ListTeamChannelSharedWithTeamsCompleteResult, error) {
	return c.ListTeamChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, options, SharedWithChannelTeamInfoOperationPredicate{})
}

// ListTeamChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelSharedWithTeamClient) ListTeamChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelSharedWithTeamsOperationOptions, predicate SharedWithChannelTeamInfoOperationPredicate) (result ListTeamChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]beta.SharedWithChannelTeamInfo, 0)

	resp, err := c.ListTeamChannelSharedWithTeams(ctx, id, options)
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

	result = ListTeamChannelSharedWithTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
