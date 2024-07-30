package lifecycleworkflowworkflowversiontask

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

type ListLifecycleWorkflowWorkflowVersionTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowWorkflowVersionTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowWorkflowVersionTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowVersionTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowVersionTasks ...
func (c LifecycleWorkflowWorkflowVersionTaskClient) ListLifecycleWorkflowWorkflowVersionTasks(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) (result ListLifecycleWorkflowWorkflowVersionTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowVersionTasksCustomPager{},
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

// ListLifecycleWorkflowWorkflowVersionTasksComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowVersionTaskClient) ListLifecycleWorkflowWorkflowVersionTasksComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId) (ListLifecycleWorkflowWorkflowVersionTasksCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowVersionTasksCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowVersionTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowVersionTaskClient) ListLifecycleWorkflowWorkflowVersionTasksCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, predicate IdentityGovernanceTaskOperationPredicate) (result ListLifecycleWorkflowWorkflowVersionTasksCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTask, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowVersionTasks(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowVersionTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
