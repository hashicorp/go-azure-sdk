package plannerfavoriteplan

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPlannerFavoritePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerPlan
}

type ListPlannerFavoritePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerPlan
}

type ListPlannerFavoritePlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerFavoritePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerFavoritePlans ...
func (c PlannerFavoritePlanClient) ListPlannerFavoritePlans(ctx context.Context, id UserId) (result ListPlannerFavoritePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerFavoritePlansCustomPager{},
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
		Values *[]beta.PlannerPlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerFavoritePlansComplete retrieves all the results into a single object
func (c PlannerFavoritePlanClient) ListPlannerFavoritePlansComplete(ctx context.Context, id UserId) (ListPlannerFavoritePlansCompleteResult, error) {
	return c.ListPlannerFavoritePlansCompleteMatchingPredicate(ctx, id, PlannerPlanOperationPredicate{})
}

// ListPlannerFavoritePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerFavoritePlanClient) ListPlannerFavoritePlansCompleteMatchingPredicate(ctx context.Context, id UserId, predicate PlannerPlanOperationPredicate) (result ListPlannerFavoritePlansCompleteResult, err error) {
	items := make([]beta.PlannerPlan, 0)

	resp, err := c.ListPlannerFavoritePlans(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListPlannerFavoritePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
