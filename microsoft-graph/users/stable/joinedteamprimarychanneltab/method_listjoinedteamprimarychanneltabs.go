package joinedteamprimarychanneltab

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

type ListJoinedTeamPrimaryChannelTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsTab
}

type ListJoinedTeamPrimaryChannelTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsTab
}

type ListJoinedTeamPrimaryChannelTabsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListJoinedTeamPrimaryChannelTabsOperationOptions() ListJoinedTeamPrimaryChannelTabsOperationOptions {
	return ListJoinedTeamPrimaryChannelTabsOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelTabsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelTabsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelTabsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelTabs - Get tabs from users. A collection of all the tabs in the channel. A navigation
// property.
func (c JoinedTeamPrimaryChannelTabClient) ListJoinedTeamPrimaryChannelTabs(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelTabsOperationOptions) (result ListJoinedTeamPrimaryChannelTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelTabsCustomPager{},
		Path:          fmt.Sprintf("%s/primaryChannel/tabs", id.ID()),
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
		Values *[]stable.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamPrimaryChannelTabsComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelTabClient) ListJoinedTeamPrimaryChannelTabsComplete(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelTabsOperationOptions) (ListJoinedTeamPrimaryChannelTabsCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx, id, options, TeamsTabOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelTabClient) ListJoinedTeamPrimaryChannelTabsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamPrimaryChannelTabsOperationOptions, predicate TeamsTabOperationPredicate) (result ListJoinedTeamPrimaryChannelTabsCompleteResult, err error) {
	items := make([]stable.TeamsTab, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelTabs(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
