package userjoinedteamscheduleopenshift

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

type ListUserByIdJoinedTeamByIdScheduleOpenShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OpenShiftCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteResult struct {
	Items []models.OpenShiftCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleOpenShifts ...
func (c UserJoinedTeamScheduleOpenShiftClient) ListUserByIdJoinedTeamByIdScheduleOpenShifts(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleOpenShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/openShifts", id.ID()),
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
		Values *[]models.OpenShiftCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleOpenShiftsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleOpenShiftClient) ListUserByIdJoinedTeamByIdScheduleOpenShiftsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteMatchingPredicate(ctx, id, models.OpenShiftCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleOpenShiftClient) ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.OpenShiftCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteResult, err error) {
	items := make([]models.OpenShiftCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleOpenShifts(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdScheduleOpenShiftsCompleteResult{
		Items: items,
	}
	return
}
