package groupteamscheduletimesoff

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

type ListGroupByIdTeamScheduleTimesOffsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TimeOffCollectionResponse
}

type ListGroupByIdTeamScheduleTimesOffsCompleteResult struct {
	Items []models.TimeOffCollectionResponse
}

// ListGroupByIdTeamScheduleTimesOffs ...
func (c GroupTeamScheduleTimesOffClient) ListGroupByIdTeamScheduleTimesOffs(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleTimesOffsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/timesOff", id.ID()),
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
		Values *[]models.TimeOffCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTeamScheduleTimesOffsComplete retrieves all the results into a single object
func (c GroupTeamScheduleTimesOffClient) ListGroupByIdTeamScheduleTimesOffsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleTimesOffsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleTimesOffsCompleteMatchingPredicate(ctx, id, models.TimeOffCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleTimesOffsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleTimesOffClient) ListGroupByIdTeamScheduleTimesOffsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.TimeOffCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleTimesOffsCompleteResult, err error) {
	items := make([]models.TimeOffCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleTimesOffs(ctx, id)
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

	result = ListGroupByIdTeamScheduleTimesOffsCompleteResult{
		Items: items,
	}
	return
}
