package joinedteamtag

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

type ListJoinedTeamTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamworkTag
}

type ListJoinedTeamTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamworkTag
}

type ListJoinedTeamTagsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListJoinedTeamTagsOperationOptions() ListJoinedTeamTagsOperationOptions {
	return ListJoinedTeamTagsOperationOptions{}
}

func (o ListJoinedTeamTagsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamTagsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamTagsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamTags - Get tags from me. The tags associated with the team.
func (c JoinedTeamTagClient) ListJoinedTeamTags(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamTagsOperationOptions) (result ListJoinedTeamTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamTagsCustomPager{},
		Path:          fmt.Sprintf("%s/tags", id.ID()),
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
		Values *[]stable.TeamworkTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamTagsComplete retrieves all the results into a single object
func (c JoinedTeamTagClient) ListJoinedTeamTagsComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamTagsOperationOptions) (ListJoinedTeamTagsCompleteResult, error) {
	return c.ListJoinedTeamTagsCompleteMatchingPredicate(ctx, id, options, TeamworkTagOperationPredicate{})
}

// ListJoinedTeamTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamTagClient) ListJoinedTeamTagsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamTagsOperationOptions, predicate TeamworkTagOperationPredicate) (result ListJoinedTeamTagsCompleteResult, err error) {
	items := make([]stable.TeamworkTag, 0)

	resp, err := c.ListJoinedTeamTags(ctx, id, options)
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

	result = ListJoinedTeamTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
