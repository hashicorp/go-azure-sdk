package mejoinedteamscheduleshift

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

type ListMeJoinedTeamByIdScheduleShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ShiftCollectionResponse
}

type ListMeJoinedTeamByIdScheduleShiftsCompleteResult struct {
	Items []models.ShiftCollectionResponse
}

// ListMeJoinedTeamByIdScheduleShifts ...
func (c MeJoinedTeamScheduleShiftClient) ListMeJoinedTeamByIdScheduleShifts(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdScheduleShiftsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdScheduleShiftsComplete retrieves all the results into a single object
func (c MeJoinedTeamScheduleShiftClient) ListMeJoinedTeamByIdScheduleShiftsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdScheduleShiftsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate(ctx, id, models.ShiftCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamScheduleShiftClient) ListMeJoinedTeamByIdScheduleShiftsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.ShiftCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdScheduleShiftsCompleteResult, err error) {
	items := make([]models.ShiftCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdScheduleShifts(ctx, id)
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

	result = ListMeJoinedTeamByIdScheduleShiftsCompleteResult{
		Items: items,
	}
	return
}
