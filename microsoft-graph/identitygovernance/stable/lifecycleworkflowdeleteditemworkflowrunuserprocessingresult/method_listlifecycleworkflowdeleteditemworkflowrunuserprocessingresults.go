package lifecycleworkflowdeleteditemworkflowrunuserprocessingresult

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

type ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResults ...
func (c LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) (result ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCustomPager{},
		Path:       fmt.Sprintf("%s/userProcessingResults", id.ID()),
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

// ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId) (ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
