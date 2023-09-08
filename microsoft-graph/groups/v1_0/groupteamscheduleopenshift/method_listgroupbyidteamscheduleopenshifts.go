package groupteamscheduleopenshift

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

type ListGroupByIdTeamScheduleOpenShiftsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OpenShiftCollectionResponse
}

type ListGroupByIdTeamScheduleOpenShiftsCompleteResult struct {
	Items []models.OpenShiftCollectionResponse
}

// ListGroupByIdTeamScheduleOpenShifts ...
func (c GroupTeamScheduleOpenShiftClient) ListGroupByIdTeamScheduleOpenShifts(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleOpenShiftsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/openShifts", id.ID()),
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

// ListGroupByIdTeamScheduleOpenShiftsComplete retrieves all the results into a single object
func (c GroupTeamScheduleOpenShiftClient) ListGroupByIdTeamScheduleOpenShiftsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleOpenShiftsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx, id, models.OpenShiftCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleOpenShiftsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleOpenShiftClient) ListGroupByIdTeamScheduleOpenShiftsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.OpenShiftCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleOpenShiftsCompleteResult, err error) {
	items := make([]models.OpenShiftCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleOpenShifts(ctx, id)
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

	result = ListGroupByIdTeamScheduleOpenShiftsCompleteResult{
		Items: items,
	}
	return
}
