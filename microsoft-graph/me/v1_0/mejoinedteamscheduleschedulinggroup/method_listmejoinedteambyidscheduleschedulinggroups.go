package mejoinedteamscheduleschedulinggroup

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

type ListMeJoinedTeamByIdScheduleSchedulingGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SchedulingGroupCollectionResponse
}

type ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteResult struct {
	Items []models.SchedulingGroupCollectionResponse
}

// ListMeJoinedTeamByIdScheduleSchedulingGroups ...
func (c MeJoinedTeamScheduleSchedulingGroupClient) ListMeJoinedTeamByIdScheduleSchedulingGroups(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdScheduleSchedulingGroupsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdScheduleSchedulingGroupsComplete retrieves all the results into a single object
func (c MeJoinedTeamScheduleSchedulingGroupClient) ListMeJoinedTeamByIdScheduleSchedulingGroupsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate(ctx, id, models.SchedulingGroupCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamScheduleSchedulingGroupClient) ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.SchedulingGroupCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteResult, err error) {
	items := make([]models.SchedulingGroupCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdScheduleSchedulingGroups(ctx, id)
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

	result = ListMeJoinedTeamByIdScheduleSchedulingGroupsCompleteResult{
		Items: items,
	}
	return
}
