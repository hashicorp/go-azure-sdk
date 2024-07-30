package lifecycleworkflowworkflowexecutionscope

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

type ListLifecycleWorkflowWorkflowExecutionScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowWorkflowExecutionScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowWorkflowExecutionScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowExecutionScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowExecutionScopes ...
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowWorkflowExecutionScopes(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (result ListLifecycleWorkflowWorkflowExecutionScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowExecutionScopesCustomPager{},
		Path:       fmt.Sprintf("%s/executionScope", id.ID()),
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
		Values *[]stable.IdentityGovernanceUserProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowWorkflowExecutionScopesComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowWorkflowExecutionScopesComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (ListLifecycleWorkflowWorkflowExecutionScopesCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowExecutionScopesCompleteMatchingPredicate(ctx, id, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowExecutionScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowWorkflowExecutionScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowWorkflowExecutionScopesCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowExecutionScopes(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowExecutionScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
