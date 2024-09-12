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

type ListPlannerPlanBucketTasksOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPlannerPlanBucketTasksOperationOptions() ListPlannerPlanBucketTasksOperationOptions {
	return ListPlannerPlanBucketTasksOperationOptions{}
}

func (o ListPlannerPlanBucketTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPlannerPlanBucketTasksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListPlannerPlanBucketTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListPlannerPlanBucketTasks - Get tasks from groups. Read-only. Nullable. The collection of tasks in the bucket.
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasks(ctx context.Context, id stable.GroupIdPlannerPlanIdBucketId, options ListPlannerPlanBucketTasksOperationOptions) (result ListPlannerPlanBucketTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPlannerPlanBucketTasksCustomPager{},
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
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
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksComplete(ctx context.Context, id stable.GroupIdPlannerPlanIdBucketId, options ListPlannerPlanBucketTasksOperationOptions) (ListPlannerPlanBucketTasksCompleteResult, error) {
	return c.ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx, id, options, PlannerTaskOperationPredicate{})
}

// ListPlannerPlanBucketTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanBucketTaskClient) ListPlannerPlanBucketTasksCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdPlannerPlanIdBucketId, options ListPlannerPlanBucketTasksOperationOptions, predicate PlannerTaskOperationPredicate) (result ListPlannerPlanBucketTasksCompleteResult, err error) {
	items := make([]stable.PlannerTask, 0)

	resp, err := c.ListPlannerPlanBucketTasks(ctx, id, options)
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
