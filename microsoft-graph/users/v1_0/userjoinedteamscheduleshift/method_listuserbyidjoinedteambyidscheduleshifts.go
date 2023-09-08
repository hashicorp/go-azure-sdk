package userjoinedteamscheduleshift

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

type ListUserByIdJoinedTeamByIdScheduleShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ShiftCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleShiftsCompleteResult struct {
	Items []models.ShiftCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleShifts ...
func (c UserJoinedTeamScheduleShiftClient) ListUserByIdJoinedTeamByIdScheduleShifts(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/shifts", id.ID()),
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
		Values *[]models.ShiftCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleShiftsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleShiftClient) ListUserByIdJoinedTeamByIdScheduleShiftsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleShiftsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate(ctx, id, models.ShiftCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleShiftClient) ListUserByIdJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.ShiftCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleShiftsCompleteResult, err error) {
	items := make([]models.ShiftCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleShifts(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdScheduleShiftsCompleteResult{
		Items: items,
	}
	return
}
