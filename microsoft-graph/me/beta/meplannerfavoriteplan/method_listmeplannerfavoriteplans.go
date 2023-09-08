package meplannerfavoriteplan

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

type ListMePlannerFavoritePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerPlanCollectionResponse
}

type ListMePlannerFavoritePlansCompleteResult struct {
	Items []models.PlannerPlanCollectionResponse
}

// ListMePlannerFavoritePlans ...
func (c MePlannerFavoritePlanClient) ListMePlannerFavoritePlans(ctx context.Context) (result ListMePlannerFavoritePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/planner/favoritePlans",
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

// ListMePlannerFavoritePlansComplete retrieves all the results into a single object
func (c MePlannerFavoritePlanClient) ListMePlannerFavoritePlansComplete(ctx context.Context) (ListMePlannerFavoritePlansCompleteResult, error) {
	return c.ListMePlannerFavoritePlansCompleteMatchingPredicate(ctx, models.PlannerPlanCollectionResponseOperationPredicate{})
}

// ListMePlannerFavoritePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerFavoritePlanClient) ListMePlannerFavoritePlansCompleteMatchingPredicate(ctx context.Context, predicate models.PlannerPlanCollectionResponseOperationPredicate) (result ListMePlannerFavoritePlansCompleteResult, err error) {
	items := make([]models.PlannerPlanCollectionResponse, 0)

	resp, err := c.ListMePlannerFavoritePlans(ctx)
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

	result = ListMePlannerFavoritePlansCompleteResult{
		Items: items,
	}
	return
}
