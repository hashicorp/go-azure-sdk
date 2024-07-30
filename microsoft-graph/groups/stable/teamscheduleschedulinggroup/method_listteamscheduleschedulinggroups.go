package teamscheduleschedulinggroup

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

type ListTeamScheduleSchedulingGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SchedulingGroup
}

type ListTeamScheduleSchedulingGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SchedulingGroup
}

type ListTeamScheduleSchedulingGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleSchedulingGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleSchedulingGroups ...
func (c TeamScheduleSchedulingGroupClient) ListTeamScheduleSchedulingGroups(ctx context.Context, id GroupId) (result ListTeamScheduleSchedulingGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleSchedulingGroupsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/schedulingGroups", id.ID()),
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

// ListTeamScheduleSchedulingGroupsComplete retrieves all the results into a single object
func (c TeamScheduleSchedulingGroupClient) ListTeamScheduleSchedulingGroupsComplete(ctx context.Context, id GroupId) (ListTeamScheduleSchedulingGroupsCompleteResult, error) {
	return c.ListTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx, id, SchedulingGroupOperationPredicate{})
}

// ListTeamScheduleSchedulingGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleSchedulingGroupClient) ListTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate SchedulingGroupOperationPredicate) (result ListTeamScheduleSchedulingGroupsCompleteResult, err error) {
	items := make([]stable.SchedulingGroup, 0)

	resp, err := c.ListTeamScheduleSchedulingGroups(ctx, id)
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

	result = ListTeamScheduleSchedulingGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
