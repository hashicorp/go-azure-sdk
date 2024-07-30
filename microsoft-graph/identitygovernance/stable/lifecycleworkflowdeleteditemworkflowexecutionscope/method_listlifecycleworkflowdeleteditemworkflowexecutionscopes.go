package lifecycleworkflowdeleteditemworkflowexecutionscope

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

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes ...
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager{},
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

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate(ctx, id, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
