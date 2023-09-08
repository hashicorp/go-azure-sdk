package groupteamscheduleoffershiftrequest

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

type ListGroupByIdTeamScheduleOfferShiftRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OfferShiftRequestCollectionResponse
}

type ListGroupByIdTeamScheduleOfferShiftRequestsCompleteResult struct {
	Items []models.OfferShiftRequestCollectionResponse
}

// ListGroupByIdTeamScheduleOfferShiftRequests ...
func (c GroupTeamScheduleOfferShiftRequestClient) ListGroupByIdTeamScheduleOfferShiftRequests(ctx context.Context, id GroupId) (result ListGroupByIdTeamScheduleOfferShiftRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/team/schedule/offerShiftRequests", id.ID()),
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

// ListGroupByIdTeamScheduleOfferShiftRequestsComplete retrieves all the results into a single object
func (c GroupTeamScheduleOfferShiftRequestClient) ListGroupByIdTeamScheduleOfferShiftRequestsComplete(ctx context.Context, id GroupId) (ListGroupByIdTeamScheduleOfferShiftRequestsCompleteResult, error) {
	return c.ListGroupByIdTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx, id, models.OfferShiftRequestCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamScheduleOfferShiftRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamScheduleOfferShiftRequestClient) ListGroupByIdTeamScheduleOfferShiftRequestsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.OfferShiftRequestCollectionResponseOperationPredicate) (result ListGroupByIdTeamScheduleOfferShiftRequestsCompleteResult, err error) {
	items := make([]models.OfferShiftRequestCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamScheduleOfferShiftRequests(ctx, id)
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

	result = ListGroupByIdTeamScheduleOfferShiftRequestsCompleteResult{
		Items: items,
	}
	return
}
