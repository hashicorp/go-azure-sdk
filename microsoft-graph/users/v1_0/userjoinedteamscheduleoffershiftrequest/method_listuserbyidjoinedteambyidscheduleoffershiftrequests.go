package userjoinedteamscheduleoffershiftrequest

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

type ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OfferShiftRequestCollectionResponse
}

type ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteResult struct {
	Items []models.OfferShiftRequestCollectionResponse
}

// ListUserByIdJoinedTeamByIdScheduleOfferShiftRequests ...
func (c UserJoinedTeamScheduleOfferShiftRequestClient) ListUserByIdJoinedTeamByIdScheduleOfferShiftRequests(ctx context.Context, id UserJoinedTeamId) (result ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/schedule/offerShiftRequests", id.ID()),
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
		Values *[]models.OfferShiftRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsComplete retrieves all the results into a single object
func (c UserJoinedTeamScheduleOfferShiftRequestClient) ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsComplete(ctx context.Context, id UserJoinedTeamId) (ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx, id, models.OfferShiftRequestCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamScheduleOfferShiftRequestClient) ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamId, predicate models.OfferShiftRequestCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteResult, err error) {
	items := make([]models.OfferShiftRequestCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdScheduleOfferShiftRequests(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdScheduleOfferShiftRequestsCompleteResult{
		Items: items,
	}
	return
}
