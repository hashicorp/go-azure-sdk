package groupteamscheduleschedulinggroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdTeamScheduleSchedulingGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SchedulingGroupCollectionResponse
}

type ListGroupByIdTeamScheduleSchedulingGroupsCompleteResult struct {
	Items []models.SchedulingGroupCollectionResponse
}

// ListGroupByIdTeamScheduleSchedulingGroups ...
func (c GroupTeamScheduleSchedulingGroupClient) ListGroupByIdTeamScheduleSchedulingGroups(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleSchedulingGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/schedulingGroups", id.ID()),
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

// ListGroupByIdTeamScheduleSchedulingGroupsComplete retrieves all the results into a single object
func (c GroupTeamScheduleSchedulingGroupClient) ListGroupByIdTeamScheduleSchedulingGroupsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleSchedulingGroupsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx, id, models.SchedulingGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleSchedulingGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleSchedulingGroupClient) ListGroupByIdTeamScheduleSchedulingGroupsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.SchedulingGroupCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleSchedulingGroupsCompleteResult, err error) {
	items := make([]models.SchedulingGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleSchedulingGroups(ctx, id)
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

	result = ListGroupByIdTeamScheduleSchedulingGroupsCompleteResult{
		Items: items,
	}
	return
}
