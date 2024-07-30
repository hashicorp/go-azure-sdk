package lifecycleworkflowworkflowtask

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

type ListLifecycleWorkflowWorkflowTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowWorkflowTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowWorkflowTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTasks ...
func (c LifecycleWorkflowWorkflowTaskClient) ListLifecycleWorkflowWorkflowTasks(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (result ListLifecycleWorkflowWorkflowTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTasksCustomPager{},
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

// ListLifecycleWorkflowWorkflowTasksComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskClient) ListLifecycleWorkflowWorkflowTasksComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (ListLifecycleWorkflowWorkflowTasksCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTasksCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskClient) ListLifecycleWorkflowWorkflowTasksCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId, predicate IdentityGovernanceTaskOperationPredicate) (result ListLifecycleWorkflowWorkflowTasksCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTask, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTasks(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
