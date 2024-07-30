package lifecycleworkflowdeleteditemworkflowtask

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

type ListLifecycleWorkflowDeletedItemWorkflowTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowDeletedItemWorkflowTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowTasks ...
func (c LifecycleWorkflowDeletedItemWorkflowTaskClient) ListLifecycleWorkflowDeletedItemWorkflowTasks(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result ListLifecycleWorkflowDeletedItemWorkflowTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowTasksCustomPager{},
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
		Values *[]stable.IdentityGovernanceTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowTasksComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowTaskClient) ListLifecycleWorkflowDeletedItemWorkflowTasksComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowTaskClient) ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, predicate IdentityGovernanceTaskOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTask, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowTasks(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
