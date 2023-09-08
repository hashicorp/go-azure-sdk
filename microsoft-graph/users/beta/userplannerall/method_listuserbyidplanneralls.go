package userplannerall

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

type ListUserByIdPlannerAllsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerDeltaCollectionResponse
}

type ListUserByIdPlannerAllsCompleteResult struct {
	Items []models.PlannerDeltaCollectionResponse
}

// ListUserByIdPlannerAlls ...
func (c UserPlannerAllClient) ListUserByIdPlannerAlls(ctx context.Context, id UserId) (result ListUserByIdPlannerAllsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/planner/all", id.ID()),
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
		Values *[]models.PlannerDeltaCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPlannerAllsComplete retrieves all the results into a single object
func (c UserPlannerAllClient) ListUserByIdPlannerAllsComplete(ctx context.Context, id UserId) (ListUserByIdPlannerAllsCompleteResult, error) {
	return c.ListUserByIdPlannerAllsCompleteMatchingPredicate(ctx, id, models.PlannerDeltaCollectionResponseOperationPredicate{})
}

// ListUserByIdPlannerAllsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPlannerAllClient) ListUserByIdPlannerAllsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.PlannerDeltaCollectionResponseOperationPredicate) (result ListUserByIdPlannerAllsCompleteResult, err error) {
	items := make([]models.PlannerDeltaCollectionResponse, 0)

	resp, err := c.ListUserByIdPlannerAlls(ctx, id)
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

	result = ListUserByIdPlannerAllsCompleteResult{
		Items: items,
	}
	return
}
