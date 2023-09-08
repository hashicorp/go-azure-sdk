package userplannerplan

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

type ListUserByIdPlannerPlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerPlanCollectionResponse
}

type ListUserByIdPlannerPlansCompleteResult struct {
	Items []models.PlannerPlanCollectionResponse
}

// ListUserByIdPlannerPlans ...
func (c UserPlannerPlanClient) ListUserByIdPlannerPlans(ctx context.Context, id UserId) (result ListUserByIdPlannerPlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/planner/plans", id.ID()),
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

// ListUserByIdPlannerPlansComplete retrieves all the results into a single object
func (c UserPlannerPlanClient) ListUserByIdPlannerPlansComplete(ctx context.Context, id UserId) (ListUserByIdPlannerPlansCompleteResult, error) {
	return c.ListUserByIdPlannerPlansCompleteMatchingPredicate(ctx, id, models.PlannerPlanCollectionResponseOperationPredicate{})
}

// ListUserByIdPlannerPlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPlannerPlanClient) ListUserByIdPlannerPlansCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PlannerPlanCollectionResponseOperationPredicate) (result ListUserByIdPlannerPlansCompleteResult, err error) {
	items := make([]models.PlannerPlanCollectionResponse, 0)

	resp, err := c.ListUserByIdPlannerPlans(ctx, id)
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

	result = ListUserByIdPlannerPlansCompleteResult{
		Items: items,
	}
	return
}
