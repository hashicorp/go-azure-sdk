package mejoinedteamscheduleopenshiftchangerequest

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

type ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OpenShiftChangeRequestCollectionResponse
}

type ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult struct {
	Items []models.OpenShiftChangeRequestCollectionResponse
}

// ListMeJoinedTeamByIdScheduleOpenShiftChangeRequests ...
func (c MeJoinedTeamScheduleOpenShiftChangeRequestClient) ListMeJoinedTeamByIdScheduleOpenShiftChangeRequests(ctx context.Context, id MeJoinedTeamId) (result ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsOperationResponse, err error) {
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

// ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsComplete retrieves all the results into a single object
func (c MeJoinedTeamScheduleOpenShiftChangeRequestClient) ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsComplete(ctx context.Context, id MeJoinedTeamId) (ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx, id, models.OpenShiftChangeRequestCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamScheduleOpenShiftChangeRequestClient) ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamId, predicate models.OpenShiftChangeRequestCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult, err error) {
	items := make([]models.OpenShiftChangeRequestCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdScheduleOpenShiftChangeRequests(ctx, id)
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

	result = ListMeJoinedTeamByIdScheduleOpenShiftChangeRequestsCompleteResult{
		Items: items,
	}
	return
}
