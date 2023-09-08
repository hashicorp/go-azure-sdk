package mejoinedteamscheduletimeoffreason

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

type ListMeJoinedTeamByIdScheduleTimeOffReasonsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TimeOffReasonCollectionResponse
}

type ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteResult struct {
	Items []models.TimeOffReasonCollectionResponse
}

// ListMeJoinedTeamByIdScheduleTimeOffReasons ...
func (c MeJoinedTeamScheduleTimeOffReasonClient) ListMeJoinedTeamByIdScheduleTimeOffReasons(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdScheduleTimeOffReasonsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/timeOffReasons", id.ID()),
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

// ListMeJoinedTeamByIdScheduleTimeOffReasonsComplete retrieves all the results into a single object
func (c MeJoinedTeamScheduleTimeOffReasonClient) ListMeJoinedTeamByIdScheduleTimeOffReasonsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteMatchingPredicate(ctx, id, models.TimeOffReasonCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamScheduleTimeOffReasonClient) ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.TimeOffReasonCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteResult, err error) {
	items := make([]models.TimeOffReasonCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdScheduleTimeOffReasons(ctx, id)
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

	result = ListMeJoinedTeamByIdScheduleTimeOffReasonsCompleteResult{
		Items: items,
	}
	return
}
