package meplannerall

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

type ListMePlannerAllsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.PlannerDeltaCollectionResponse
}

type ListMePlannerAllsCompleteResult struct {
	Items []models.PlannerDeltaCollectionResponse
}

// ListMePlannerAlls ...
func (c MePlannerAllClient) ListMePlannerAlls(ctx context.Context) (result ListMePlannerAllsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/planner/all",
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

// ListMePlannerAllsComplete retrieves all the results into a single object
func (c MePlannerAllClient) ListMePlannerAllsComplete(ctx context.Context) (ListMePlannerAllsCompleteResult, error) {
	return c.ListMePlannerAllsCompleteMatchingPredicate(ctx, models.PlannerDeltaCollectionResponseOperationPredicate{})
}

// ListMePlannerAllsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MePlannerAllClient) ListMePlannerAllsCompleteMatchingPredicate(ctx context.Context, predicate models.PlannerDeltaCollectionResponseOperationPredicate) (result ListMePlannerAllsCompleteResult, err error) {
	items := make([]models.PlannerDeltaCollectionResponse, 0)

	resp, err := c.ListMePlannerAlls(ctx)
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

	result = ListMePlannerAllsCompleteResult{
		Items: items,
	}
	return
}
