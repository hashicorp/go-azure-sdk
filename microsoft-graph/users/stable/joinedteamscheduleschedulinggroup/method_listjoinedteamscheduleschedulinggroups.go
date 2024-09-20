package joinedteamscheduleschedulinggroup

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

type ListJoinedTeamScheduleSchedulingGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SchedulingGroup
}

type ListJoinedTeamScheduleSchedulingGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SchedulingGroup
}

type ListJoinedTeamScheduleSchedulingGroupsOperationOptions struct {
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

func DefaultListJoinedTeamScheduleSchedulingGroupsOperationOptions() ListJoinedTeamScheduleSchedulingGroupsOperationOptions {
	return ListJoinedTeamScheduleSchedulingGroupsOperationOptions{}
}

func (o ListJoinedTeamScheduleSchedulingGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamScheduleSchedulingGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamScheduleSchedulingGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamScheduleSchedulingGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleSchedulingGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleSchedulingGroups - Get schedulingGroups from users. The logical grouping of users in the
// schedule (usually by role).
func (c JoinedTeamScheduleSchedulingGroupClient) ListJoinedTeamScheduleSchedulingGroups(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamScheduleSchedulingGroupsOperationOptions) (result ListJoinedTeamScheduleSchedulingGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamScheduleSchedulingGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/schedule/schedulingGroups", id.ID()),
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
		Values *[]stable.SchedulingGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleSchedulingGroupsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleSchedulingGroupClient) ListJoinedTeamScheduleSchedulingGroupsComplete(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamScheduleSchedulingGroupsOperationOptions) (ListJoinedTeamScheduleSchedulingGroupsCompleteResult, error) {
	return c.ListJoinedTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx, id, options, SchedulingGroupOperationPredicate{})
}

// ListJoinedTeamScheduleSchedulingGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleSchedulingGroupClient) ListJoinedTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamId, options ListJoinedTeamScheduleSchedulingGroupsOperationOptions, predicate SchedulingGroupOperationPredicate) (result ListJoinedTeamScheduleSchedulingGroupsCompleteResult, err error) {
	items := make([]stable.SchedulingGroup, 0)

	resp, err := c.ListJoinedTeamScheduleSchedulingGroups(ctx, id, options)
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

	result = ListJoinedTeamScheduleSchedulingGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
