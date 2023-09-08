package groupteamscheduletimeoffreason

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

type ListGroupByIdTeamScheduleTimeOffReasonsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TimeOffReasonCollectionResponse
}

type ListGroupByIdTeamScheduleTimeOffReasonsCompleteResult struct {
	Items []models.TimeOffReasonCollectionResponse
}

// ListGroupByIdTeamScheduleTimeOffReasons ...
func (c GroupTeamScheduleTimeOffReasonClient) ListGroupByIdTeamScheduleTimeOffReasons(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleTimeOffReasonsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/timeOffReasons", id.ID()),
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
		Values *[]models.TimeOffReasonCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamScheduleTimeOffReasonsComplete retrieves all the results into a single object
func (c GroupTeamScheduleTimeOffReasonClient) ListGroupByIdTeamScheduleTimeOffReasonsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleTimeOffReasonsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleTimeOffReasonsCompleteMatchingPredicate(ctx, id, models.TimeOffReasonCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleTimeOffReasonsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleTimeOffReasonClient) ListGroupByIdTeamScheduleTimeOffReasonsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.TimeOffReasonCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleTimeOffReasonsCompleteResult, err error) {
	items := make([]models.TimeOffReasonCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleTimeOffReasons(ctx, id)
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

	result = ListGroupByIdTeamScheduleTimeOffReasonsCompleteResult{
		Items: items,
	}
	return
}
