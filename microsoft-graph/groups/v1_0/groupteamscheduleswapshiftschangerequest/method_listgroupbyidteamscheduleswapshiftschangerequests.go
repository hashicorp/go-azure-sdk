package groupteamscheduleswapshiftschangerequest

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

type ListGroupByIdTeamScheduleSwapShiftsChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SwapShiftsChangeRequestCollectionResponse
}

type ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteResult struct {
	Items []models.SwapShiftsChangeRequestCollectionResponse
}

// ListGroupByIdTeamScheduleSwapShiftsChangeRequests ...
func (c GroupTeamScheduleSwapShiftsChangeRequestClient) ListGroupByIdTeamScheduleSwapShiftsChangeRequests(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleSwapShiftsChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.SwapShiftsChangeRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamScheduleSwapShiftsChangeRequestsComplete retrieves all the results into a single object
func (c GroupTeamScheduleSwapShiftsChangeRequestClient) ListGroupByIdTeamScheduleSwapShiftsChangeRequestsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx, id, models.SwapShiftsChangeRequestCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleSwapShiftsChangeRequestClient) ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.SwapShiftsChangeRequestCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteResult, err error) {
	items := make([]models.SwapShiftsChangeRequestCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleSwapShiftsChangeRequests(ctx, id)
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

	result = ListGroupByIdTeamScheduleSwapShiftsChangeRequestsCompleteResult{
		Items: items,
	}
	return
}
