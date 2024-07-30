package joinedteamscheduleswapshiftschangerequest

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

type ListJoinedTeamScheduleSwapShiftsChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SwapShiftsChangeRequest
}

type ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SwapShiftsChangeRequest
}

type ListJoinedTeamScheduleSwapShiftsChangeRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamScheduleSwapShiftsChangeRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamScheduleSwapShiftsChangeRequests ...
func (c JoinedTeamScheduleSwapShiftsChangeRequestClient) ListJoinedTeamScheduleSwapShiftsChangeRequests(ctx context.Context, id MeJoinedTeamId) (result ListJoinedTeamScheduleSwapShiftsChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamScheduleSwapShiftsChangeRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/schedule/swapShiftsChangeRequests", id.ID()),
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
		Values *[]stable.SwapShiftsChangeRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamScheduleSwapShiftsChangeRequestsComplete retrieves all the results into a single object
func (c JoinedTeamScheduleSwapShiftsChangeRequestClient) ListJoinedTeamScheduleSwapShiftsChangeRequestsComplete(ctx context.Context, id MeJoinedTeamId) (ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteResult, error) {
	return c.ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx, id, SwapShiftsChangeRequestOperationPredicate{})
}

// ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamScheduleSwapShiftsChangeRequestClient) ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate SwapShiftsChangeRequestOperationPredicate) (result ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteResult, err error) {
	items := make([]stable.SwapShiftsChangeRequest, 0)

	resp, err := c.ListJoinedTeamScheduleSwapShiftsChangeRequests(ctx, id)
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

	result = ListJoinedTeamScheduleSwapShiftsChangeRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
