package plannerrosterplan

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

type ListPlannerRosterPlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerPlan
}

type ListPlannerRosterPlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerPlan
}

type ListPlannerRosterPlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerRosterPlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerRosterPlans ...
func (c PlannerRosterPlanClient) ListPlannerRosterPlans(ctx context.Context) (result ListPlannerRosterPlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerRosterPlansCustomPager{},
		Path:       "/me/planner/rosterPlans",
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

// ListPlannerRosterPlansComplete retrieves all the results into a single object
func (c PlannerRosterPlanClient) ListPlannerRosterPlansComplete(ctx context.Context) (ListPlannerRosterPlansCompleteResult, error) {
	return c.ListPlannerRosterPlansCompleteMatchingPredicate(ctx, PlannerPlanOperationPredicate{})
}

// ListPlannerRosterPlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerRosterPlanClient) ListPlannerRosterPlansCompleteMatchingPredicate(ctx context.Context, predicate PlannerPlanOperationPredicate) (result ListPlannerRosterPlansCompleteResult, err error) {
	items := make([]beta.PlannerPlan, 0)

	resp, err := c.ListPlannerRosterPlans(ctx)
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

	result = ListPlannerRosterPlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
