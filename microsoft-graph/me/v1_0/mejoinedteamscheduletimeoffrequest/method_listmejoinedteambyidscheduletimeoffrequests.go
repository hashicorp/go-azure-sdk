package mejoinedteamscheduletimeoffrequest

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

type ListMeJoinedTeamByIdScheduleTimeOffRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TimeOffRequestCollectionResponse
}

type ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteResult struct {
	Items []models.TimeOffRequestCollectionResponse
}

// ListMeJoinedTeamByIdScheduleTimeOffRequests ...
func (c MeJoinedTeamScheduleTimeOffRequestClient) ListMeJoinedTeamByIdScheduleTimeOffRequests(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdScheduleTimeOffRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/timeOffRequests", id.ID()),
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
		Values *[]models.TimeOffRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeJoinedTeamByIdScheduleTimeOffRequestsComplete retrieves all the results into a single object
func (c MeJoinedTeamScheduleTimeOffRequestClient) ListMeJoinedTeamByIdScheduleTimeOffRequestsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteMatchingPredicate(ctx, id, models.TimeOffRequestCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamScheduleTimeOffRequestClient) ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.TimeOffRequestCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteResult, err error) {
	items := make([]models.TimeOffRequestCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdScheduleTimeOffRequests(ctx, id)
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

	result = ListMeJoinedTeamByIdScheduleTimeOffRequestsCompleteResult{
		Items: items,
	}
	return
}
