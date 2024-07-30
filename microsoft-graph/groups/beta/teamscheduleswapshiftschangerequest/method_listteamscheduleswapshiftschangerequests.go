package teamscheduleswapshiftschangerequest

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

type ListTeamScheduleSwapShiftsChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SwapShiftsChangeRequest
}

type ListTeamScheduleSwapShiftsChangeRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SwapShiftsChangeRequest
}

type ListTeamScheduleSwapShiftsChangeRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleSwapShiftsChangeRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleSwapShiftsChangeRequests ...
func (c TeamScheduleSwapShiftsChangeRequestClient) ListTeamScheduleSwapShiftsChangeRequests(ctx context.Context, id GroupId) (result ListTeamScheduleSwapShiftsChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamScheduleSwapShiftsChangeRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/team/schedule/swapShiftsChangeRequests", id.ID()),
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
		Values *[]beta.SwapShiftsChangeRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleSwapShiftsChangeRequestsComplete retrieves all the results into a single object
func (c TeamScheduleSwapShiftsChangeRequestClient) ListTeamScheduleSwapShiftsChangeRequestsComplete(ctx context.Context, id GroupId) (ListTeamScheduleSwapShiftsChangeRequestsCompleteResult, error) {
	return c.ListTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx, id, SwapShiftsChangeRequestOperationPredicate{})
}

// ListTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleSwapShiftsChangeRequestClient) ListTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate SwapShiftsChangeRequestOperationPredicate) (result ListTeamScheduleSwapShiftsChangeRequestsCompleteResult, err error) {
	items := make([]beta.SwapShiftsChangeRequest, 0)

	resp, err := c.ListTeamScheduleSwapShiftsChangeRequests(ctx, id)
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

	result = ListTeamScheduleSwapShiftsChangeRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
