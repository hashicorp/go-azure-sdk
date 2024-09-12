package teamtagmember

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

type ListTeamTagMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamworkTagMember
}

type ListTeamTagMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamworkTagMember
}

type ListTeamTagMembersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTeamTagMembersOperationOptions() ListTeamTagMembersOperationOptions {
	return ListTeamTagMembersOperationOptions{}
}

func (o ListTeamTagMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamTagMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamTagMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamTagMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamTagMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamTagMembers - Get members from groups. Users assigned to the tag.
func (c TeamTagMemberClient) ListTeamTagMembers(ctx context.Context, id beta.GroupIdTeamTagId, options ListTeamTagMembersOperationOptions) (result ListTeamTagMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamTagMembersCustomPager{},
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
		Values *[]beta.TeamworkTagMember `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamTagMembersComplete retrieves all the results into a single object
func (c TeamTagMemberClient) ListTeamTagMembersComplete(ctx context.Context, id beta.GroupIdTeamTagId, options ListTeamTagMembersOperationOptions) (ListTeamTagMembersCompleteResult, error) {
	return c.ListTeamTagMembersCompleteMatchingPredicate(ctx, id, options, TeamworkTagMemberOperationPredicate{})
}

// ListTeamTagMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamTagMemberClient) ListTeamTagMembersCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamTagId, options ListTeamTagMembersOperationOptions, predicate TeamworkTagMemberOperationPredicate) (result ListTeamTagMembersCompleteResult, err error) {
	items := make([]beta.TeamworkTagMember, 0)

	resp, err := c.ListTeamTagMembers(ctx, id, options)
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

	result = ListTeamTagMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
