package teamscheduletimeoffreason

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

type ListTeamScheduleTimeOffReasonsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TimeOffReason
}

type ListTeamScheduleTimeOffReasonsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TimeOffReason
}

type ListTeamScheduleTimeOffReasonsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleTimeOffReasonsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleTimeOffReasons ...
func (c TeamScheduleTimeOffReasonClient) ListTeamScheduleTimeOffReasons(ctx context.Context, id GroupId) (result ListTeamScheduleTimeOffReasonsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleTimeOffReasonsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/timeOffReasons", id.ID()),
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
		Values *[]stable.TimeOffReason `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleTimeOffReasonsComplete retrieves all the results into a single object
func (c TeamScheduleTimeOffReasonClient) ListTeamScheduleTimeOffReasonsComplete(ctx context.Context, id GroupId) (ListTeamScheduleTimeOffReasonsCompleteResult, error) {
	return c.ListTeamScheduleTimeOffReasonsCompleteMatchingPredicate(ctx, id, TimeOffReasonOperationPredicate{})
}

// ListTeamScheduleTimeOffReasonsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleTimeOffReasonClient) ListTeamScheduleTimeOffReasonsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TimeOffReasonOperationPredicate) (result ListTeamScheduleTimeOffReasonsCompleteResult, err error) {
	items := make([]stable.TimeOffReason, 0)

	resp, err := c.ListTeamScheduleTimeOffReasons(ctx, id)
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

	result = ListTeamScheduleTimeOffReasonsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
