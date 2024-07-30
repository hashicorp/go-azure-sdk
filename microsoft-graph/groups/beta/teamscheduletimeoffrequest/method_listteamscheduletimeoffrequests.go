package teamscheduletimeoffrequest

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

type ListTeamScheduleTimeOffRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TimeOffRequest
}

type ListTeamScheduleTimeOffRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TimeOffRequest
}

type ListTeamScheduleTimeOffRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleTimeOffRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleTimeOffRequests ...
func (c TeamScheduleTimeOffRequestClient) ListTeamScheduleTimeOffRequests(ctx context.Context, id GroupId) (result ListTeamScheduleTimeOffRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleTimeOffRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/timeOffRequests", id.ID()),
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
		Values *[]beta.TimeOffRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleTimeOffRequestsComplete retrieves all the results into a single object
func (c TeamScheduleTimeOffRequestClient) ListTeamScheduleTimeOffRequestsComplete(ctx context.Context, id GroupId) (ListTeamScheduleTimeOffRequestsCompleteResult, error) {
	return c.ListTeamScheduleTimeOffRequestsCompleteMatchingPredicate(ctx, id, TimeOffRequestOperationPredicate{})
}

// ListTeamScheduleTimeOffRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleTimeOffRequestClient) ListTeamScheduleTimeOffRequestsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TimeOffRequestOperationPredicate) (result ListTeamScheduleTimeOffRequestsCompleteResult, err error) {
	items := make([]beta.TimeOffRequest, 0)

	resp, err := c.ListTeamScheduleTimeOffRequests(ctx, id)
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

	result = ListTeamScheduleTimeOffRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
