package userjoinedteamscheduleschedulinggroup

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

type ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SchedulingGroupCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteResult struct {
	Items []models.SchedulingGroupCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleSchedulingGroups ...
func (c UserJoinedTeamScheduleSchedulingGroupClient) ListUserByIdJoinedTeamByIdScheduleSchedulingGroups(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/schedulingGroups", id.ID()),
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
		Values *[]models.SchedulingGroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleSchedulingGroupClient) ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate(ctx, id, models.SchedulingGroupCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleSchedulingGroupClient) ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.SchedulingGroupCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteResult, err error) {
	items := make([]models.SchedulingGroupCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleSchedulingGroups(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdScheduleSchedulingGroupsCompleteResult{
		Items: items,
	}
	return
}
