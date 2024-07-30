package lifecycleworkflowworkflow

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

type ListLifecycleWorkflowWorkflowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowWorkflowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowWorkflowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflows ...
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflows(ctx context.Context) (result ListLifecycleWorkflowWorkflowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowsCustomPager{},
		Path:       "/identityGovernance/lifecycleWorkflows/workflows",
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
		Values *[]stable.IdentityGovernanceWorkflow `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowWorkflowsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflowsComplete(ctx context.Context) (ListLifecycleWorkflowWorkflowsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate(ctx, IdentityGovernanceWorkflowOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate(ctx context.Context, predicate IdentityGovernanceWorkflowOperationPredicate) (result ListLifecycleWorkflowWorkflowsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflow, 0)

	resp, err := c.ListLifecycleWorkflowWorkflows(ctx)
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

	result = ListLifecycleWorkflowWorkflowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
