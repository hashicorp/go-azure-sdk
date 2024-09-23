package teamchanneltab

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

type ListTeamChannelTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsTab
}

type ListTeamChannelTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsTab
}

type ListTeamChannelTabsOperationOptions struct {
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

func DefaultListTeamChannelTabsOperationOptions() ListTeamChannelTabsOperationOptions {
	return ListTeamChannelTabsOperationOptions{}
}

func (o ListTeamChannelTabsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelTabsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelTabsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelTabs - Get tabs from groups. A collection of all the tabs in the channel. A navigation property.
func (c TeamChannelTabClient) ListTeamChannelTabs(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelTabsOperationOptions) (result ListTeamChannelTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelTabsCustomPager{},
		Path:          fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]stable.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelTabsComplete retrieves all the results into a single object
func (c TeamChannelTabClient) ListTeamChannelTabsComplete(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelTabsOperationOptions) (ListTeamChannelTabsCompleteResult, error) {
	return c.ListTeamChannelTabsCompleteMatchingPredicate(ctx, id, options, TeamsTabOperationPredicate{})
}

// ListTeamChannelTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelTabClient) ListTeamChannelTabsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelTabsOperationOptions, predicate TeamsTabOperationPredicate) (result ListTeamChannelTabsCompleteResult, err error) {
	items := make([]stable.TeamsTab, 0)

	resp, err := c.ListTeamChannelTabs(ctx, id, options)
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

	result = ListTeamChannelTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
