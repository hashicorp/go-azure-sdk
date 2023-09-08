package groupplannerplan

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

type ListGroupByIdPlannerPlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerPlanCollectionResponse
}

type ListGroupByIdPlannerPlansCompleteResult struct {
	Items []models.PlannerPlanCollectionResponse
}

// ListGroupByIdPlannerPlans ...
func (c GroupPlannerPlanClient) ListGroupByIdPlannerPlans(ctx context.Context, id GroupId) (result ListGroupByIdPlannerPlansOperationResponse, err error) {
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

// ListGroupByIdPlannerPlansComplete retrieves all the results into a single object
func (c GroupPlannerPlanClient) ListGroupByIdPlannerPlansComplete(ctx context.Context, id GroupId) (ListGroupByIdPlannerPlansCompleteResult, error) {
	return c.ListGroupByIdPlannerPlansCompleteMatchingPredicate(ctx, id, models.PlannerPlanCollectionResponseOperationPredicate{})
}

// ListGroupByIdPlannerPlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPlannerPlanClient) ListGroupByIdPlannerPlansCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.PlannerPlanCollectionResponseOperationPredicate) (result ListGroupByIdPlannerPlansCompleteResult, err error) {
	items := make([]models.PlannerPlanCollectionResponse, 0)

	resp, err := c.ListGroupByIdPlannerPlans(ctx, id)
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

	result = ListGroupByIdPlannerPlansCompleteResult{
		Items: items,
	}
	return
}
