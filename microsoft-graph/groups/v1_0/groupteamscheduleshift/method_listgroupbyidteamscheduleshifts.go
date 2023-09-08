package groupteamscheduleshift

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

type ListGroupByIdTeamScheduleShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ShiftCollectionResponse
}

type ListGroupByIdTeamScheduleShiftsCompleteResult struct {
	Items []models.ShiftCollectionResponse
}

// ListGroupByIdTeamScheduleShifts ...
func (c GroupTeamScheduleShiftClient) ListGroupByIdTeamScheduleShifts(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/shifts", id.ID()),
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

// ListGroupByIdTeamScheduleShiftsComplete retrieves all the results into a single object
func (c GroupTeamScheduleShiftClient) ListGroupByIdTeamScheduleShiftsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleShiftsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleShiftsCompleteMatchingPredicate(ctx, id, models.ShiftCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleShiftClient) ListGroupByIdTeamScheduleShiftsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.ShiftCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleShiftsCompleteResult, err error) {
	items := make([]models.ShiftCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleShifts(ctx, id)
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

	result = ListGroupByIdTeamScheduleShiftsCompleteResult{
		Items: items,
	}
	return
}
