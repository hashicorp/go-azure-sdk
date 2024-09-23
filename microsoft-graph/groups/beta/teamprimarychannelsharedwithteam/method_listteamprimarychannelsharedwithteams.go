package teamprimarychannelsharedwithteam

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

type ListTeamPrimaryChannelSharedWithTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SharedWithChannelTeamInfo
}

type ListTeamPrimaryChannelSharedWithTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SharedWithChannelTeamInfo
}

type ListTeamPrimaryChannelSharedWithTeamsOperationOptions struct {
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

func DefaultListTeamPrimaryChannelSharedWithTeamsOperationOptions() ListTeamPrimaryChannelSharedWithTeamsOperationOptions {
	return ListTeamPrimaryChannelSharedWithTeamsOperationOptions{}
}

func (o ListTeamPrimaryChannelSharedWithTeamsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPrimaryChannelSharedWithTeamsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPrimaryChannelSharedWithTeamsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPrimaryChannelSharedWithTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelSharedWithTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelSharedWithTeams - Get sharedWithTeams from groups. A collection of teams with which a channel
// is shared.
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeams(ctx context.Context, id beta.GroupId, options ListTeamPrimaryChannelSharedWithTeamsOperationOptions) (result ListTeamPrimaryChannelSharedWithTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPrimaryChannelSharedWithTeamsCustomPager{},
		Path:          fmt.Sprintf("%s/team/primaryChannel/sharedWithTeams", id.ID()),
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
		Values *[]beta.SharedWithChannelTeamInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelSharedWithTeamsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeamsComplete(ctx context.Context, id beta.GroupId, options ListTeamPrimaryChannelSharedWithTeamsOperationOptions) (ListTeamPrimaryChannelSharedWithTeamsCompleteResult, error) {
	return c.ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx, id, options, SharedWithChannelTeamInfoOperationPredicate{})
}

// ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelSharedWithTeamClient) ListTeamPrimaryChannelSharedWithTeamsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTeamPrimaryChannelSharedWithTeamsOperationOptions, predicate SharedWithChannelTeamInfoOperationPredicate) (result ListTeamPrimaryChannelSharedWithTeamsCompleteResult, err error) {
	items := make([]beta.SharedWithChannelTeamInfo, 0)

	resp, err := c.ListTeamPrimaryChannelSharedWithTeams(ctx, id, options)
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

	result = ListTeamPrimaryChannelSharedWithTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
