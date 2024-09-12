package joinedteamtagmember

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

type ListJoinedTeamTagMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamworkTagMember
}

type ListJoinedTeamTagMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamworkTagMember
}

type ListJoinedTeamTagMembersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListJoinedTeamTagMembersOperationOptions() ListJoinedTeamTagMembersOperationOptions {
	return ListJoinedTeamTagMembersOperationOptions{}
}

func (o ListJoinedTeamTagMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamTagMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamTagMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamTagMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamTagMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamTagMembers - Get members from me. Users assigned to the tag.
func (c JoinedTeamTagMemberClient) ListJoinedTeamTagMembers(ctx context.Context, id stable.MeJoinedTeamIdTagId, options ListJoinedTeamTagMembersOperationOptions) (result ListJoinedTeamTagMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamTagMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members", id.ID()),
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
		Values *[]stable.TeamworkTagMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamTagMembersComplete retrieves all the results into a single object
func (c JoinedTeamTagMemberClient) ListJoinedTeamTagMembersComplete(ctx context.Context, id stable.MeJoinedTeamIdTagId, options ListJoinedTeamTagMembersOperationOptions) (ListJoinedTeamTagMembersCompleteResult, error) {
	return c.ListJoinedTeamTagMembersCompleteMatchingPredicate(ctx, id, options, TeamworkTagMemberOperationPredicate{})
}

// ListJoinedTeamTagMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamTagMemberClient) ListJoinedTeamTagMembersCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamIdTagId, options ListJoinedTeamTagMembersOperationOptions, predicate TeamworkTagMemberOperationPredicate) (result ListJoinedTeamTagMembersCompleteResult, err error) {
	items := make([]stable.TeamworkTagMember, 0)

	resp, err := c.ListJoinedTeamTagMembers(ctx, id, options)
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

	result = ListJoinedTeamTagMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
