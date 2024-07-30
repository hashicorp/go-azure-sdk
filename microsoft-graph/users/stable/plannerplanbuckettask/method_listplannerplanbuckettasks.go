package plannerplanbuckettask

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

type ListPlannerPlanBucketTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PlannerTask
}

type ListPlannerPlanBucketTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PlannerTask
}

type ListPlannerPlanBucketTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerPlanBucketTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerPlanBucketTasks ...
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasks(ctx context.Context, id UserIdPlannerPlanIdBucketId) (result ListPlannerPlanBucketTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerPlanBucketTasksCustomPager{},
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

// ListPlannerPlanBucketTasksComplete retrieves all the results into a single object
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksComplete(ctx context.Context, id UserIdPlannerPlanIdBucketId) (ListPlannerPlanBucketTasksCompleteResult, error) {
	return c.ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx, id, PlannerTaskOperationPredicate{})
}

// ListPlannerPlanBucketTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx context.Context, id UserIdPlannerPlanIdBucketId, predicate PlannerTaskOperationPredicate) (result ListPlannerPlanBucketTasksCompleteResult, err error) {
	items := make([]stable.PlannerTask, 0)

	resp, err := c.ListPlannerPlanBucketTasks(ctx, id)
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

	result = ListPlannerPlanBucketTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
