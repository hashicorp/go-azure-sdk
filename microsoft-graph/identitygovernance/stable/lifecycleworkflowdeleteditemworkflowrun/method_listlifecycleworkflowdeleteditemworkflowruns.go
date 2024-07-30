package lifecycleworkflowdeleteditemworkflowrun

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

type ListLifecycleWorkflowDeletedItemWorkflowRunsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceRun
}

type ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceRun
}

type ListLifecycleWorkflowDeletedItemWorkflowRunsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowRunsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowRuns ...
func (c LifecycleWorkflowDeletedItemWorkflowRunClient) ListLifecycleWorkflowDeletedItemWorkflowRuns(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result ListLifecycleWorkflowDeletedItemWorkflowRunsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowRunsCustomPager{},
		Path:       fmt.Sprintf("%s/runs", id.ID()),
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
		Values *[]stable.IdentityGovernanceRun `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowRunsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowRunClient) ListLifecycleWorkflowDeletedItemWorkflowRunsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteMatchingPredicate(ctx, id, IdentityGovernanceRunOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowRunClient) ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, predicate IdentityGovernanceRunOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceRun, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowRuns(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowRunsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
