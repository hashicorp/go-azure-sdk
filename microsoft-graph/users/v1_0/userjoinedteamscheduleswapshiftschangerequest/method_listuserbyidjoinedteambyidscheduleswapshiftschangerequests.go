package userjoinedteamscheduleswapshiftschangerequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SwapShiftsChangeRequestCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteResult struct {
	Items []models.SwapShiftsChangeRequestCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequests ...
func (c UserJoinedTeamScheduleSwapShiftsChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequests(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.SwapShiftsChangeRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleSwapShiftsChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx, id, models.SwapShiftsChangeRequestCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleSwapShiftsChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.SwapShiftsChangeRequestCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteResult, err error) {
	items := make([]models.SwapShiftsChangeRequestCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequests(ctx, id)
	if err != nil {
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

	result = ListUserByIdJoinedTeamByIdScheduleSwapShiftsChangeRequestsCompleteResult{
		Items: items,
	}
	return
}
