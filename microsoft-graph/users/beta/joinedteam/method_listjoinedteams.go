package joinedteam

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

type ListJoinedTeamsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Team
}

type ListJoinedTeamsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Team
}

type ListJoinedTeamsOperationOptions struct {
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

func DefaultListJoinedTeamsOperationOptions() ListJoinedTeamsOperationOptions {
	return ListJoinedTeamsOperationOptions{}
}

func (o ListJoinedTeamsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeams - Get joinedTeams from users. The Microsoft Teams teams the user is a member of. Read-only. Nullable.
func (c JoinedTeamClient) ListJoinedTeams(ctx context.Context, id beta.UserId, options ListJoinedTeamsOperationOptions) (result ListJoinedTeamsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamsCustomPager{},
		Path:          fmt.Sprintf("%s/joinedTeams", id.ID()),
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
		Values *[]beta.Team `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamsComplete retrieves all the results into a single object
func (c JoinedTeamClient) ListJoinedTeamsComplete(ctx context.Context, id beta.UserId, options ListJoinedTeamsOperationOptions) (ListJoinedTeamsCompleteResult, error) {
	return c.ListJoinedTeamsCompleteMatchingPredicate(ctx, id, options, TeamOperationPredicate{})
}

// ListJoinedTeamsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamClient) ListJoinedTeamsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListJoinedTeamsOperationOptions, predicate TeamOperationPredicate) (result ListJoinedTeamsCompleteResult, err error) {
	items := make([]beta.Team, 0)

	resp, err := c.ListJoinedTeams(ctx, id, options)
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

	result = ListJoinedTeamsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
