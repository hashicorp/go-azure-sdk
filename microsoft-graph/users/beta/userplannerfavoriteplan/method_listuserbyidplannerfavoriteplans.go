package userplannerfavoriteplan

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

type ListUserByIdPlannerFavoritePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerPlanCollectionResponse
}

type ListUserByIdPlannerFavoritePlansCompleteResult struct {
	Items []models.PlannerPlanCollectionResponse
}

// ListUserByIdPlannerFavoritePlans ...
func (c UserPlannerFavoritePlanClient) ListUserByIdPlannerFavoritePlans(ctx context.Context, id UserId) (result ListUserByIdPlannerFavoritePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/planner/favoritePlans", id.ID()),
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
		Values *[]models.PlannerPlanCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPlannerFavoritePlansComplete retrieves all the results into a single object
func (c UserPlannerFavoritePlanClient) ListUserByIdPlannerFavoritePlansComplete(ctx context.Context, id UserId) (ListUserByIdPlannerFavoritePlansCompleteResult, error) {
	return c.ListUserByIdPlannerFavoritePlansCompleteMatchingPredicate(ctx, id, models.PlannerPlanCollectionResponseOperationPredicate{})
}

// ListUserByIdPlannerFavoritePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPlannerFavoritePlanClient) ListUserByIdPlannerFavoritePlansCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PlannerPlanCollectionResponseOperationPredicate) (result ListUserByIdPlannerFavoritePlansCompleteResult, err error) {
	items := make([]models.PlannerPlanCollectionResponse, 0)

	resp, err := c.ListUserByIdPlannerFavoritePlans(ctx, id)
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

	result = ListUserByIdPlannerFavoritePlansCompleteResult{
		Items: items,
	}
	return
}
