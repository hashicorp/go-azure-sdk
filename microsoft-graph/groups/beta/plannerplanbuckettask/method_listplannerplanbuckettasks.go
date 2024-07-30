package plannerplanbuckettask

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

type ListPlannerPlanBucketTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerTask
}

type ListPlannerPlanBucketTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerTask
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
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasks(ctx context.Context, id GroupIdPlannerPlanIdBucketId) (result ListPlannerPlanBucketTasksOperationResponse, err error) {
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
		Values *[]beta.PlannerTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerPlanBucketTasksComplete retrieves all the results into a single object
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksComplete(ctx context.Context, id GroupIdPlannerPlanIdBucketId) (ListPlannerPlanBucketTasksCompleteResult, error) {
	return c.ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx, id, PlannerTaskOperationPredicate{})
}

// ListPlannerPlanBucketTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx context.Context, id GroupIdPlannerPlanIdBucketId, predicate PlannerTaskOperationPredicate) (result ListPlannerPlanBucketTasksCompleteResult, err error) {
	items := make([]beta.PlannerTask, 0)

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
