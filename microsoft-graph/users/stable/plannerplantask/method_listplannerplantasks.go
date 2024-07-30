package plannerplantask

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPlannerPlanTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PlannerTask
}

type ListPlannerPlanTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PlannerTask
}

type ListPlannerPlanTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerPlanTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerPlanTasks ...
func (c PlannerPlanTaskClient) ListPlannerPlanTasks(ctx context.Context, id UserIdPlannerPlanId) (result ListPlannerPlanTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerPlanTasksCustomPager{},
		Path:       fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]stable.PlannerTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerPlanTasksComplete retrieves all the results into a single object
func (c PlannerPlanTaskClient) ListPlannerPlanTasksComplete(ctx context.Context, id UserIdPlannerPlanId) (ListPlannerPlanTasksCompleteResult, error) {
	return c.ListPlannerPlanTasksCompleteMatchingPredicate(ctx, id, PlannerTaskOperationPredicate{})
}

// ListPlannerPlanTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanTaskClient) ListPlannerPlanTasksCompleteMatchingPredicate(ctx context.Context, id UserIdPlannerPlanId, predicate PlannerTaskOperationPredicate) (result ListPlannerPlanTasksCompleteResult, err error) {
	items := make([]stable.PlannerTask, 0)

	resp, err := c.ListPlannerPlanTasks(ctx, id)
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

	result = ListPlannerPlanTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
