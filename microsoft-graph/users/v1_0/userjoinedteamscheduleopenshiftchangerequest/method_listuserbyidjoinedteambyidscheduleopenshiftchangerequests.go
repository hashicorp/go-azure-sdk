package userjoinedteamscheduleopenshiftchangerequest

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

type ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OpenShiftChangeRequestCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult struct {
	Items []models.OpenShiftChangeRequestCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequests ...
func (c UserJoinedTeamScheduleOpenShiftChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequests(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/openShiftChangeRequests", id.ID()),
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
		Values *[]models.OpenShiftChangeRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleOpenShiftChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx, id, models.OpenShiftChangeRequestCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleOpenShiftChangeRequestClient) ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.OpenShiftChangeRequestCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult, err error) {
	items := make([]models.OpenShiftChangeRequestCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequests(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult{
		Items: items,
	}
	return
}
