package lifecycleworkflowworkflowrunuserprocessingresult

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

type ListLifecycleWorkflowWorkflowRunUserProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowWorkflowRunUserProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowRunUserProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowRunUserProcessingResults ...
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowWorkflowRunUserProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) (result ListLifecycleWorkflowWorkflowRunUserProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowRunUserProcessingResultsCustomPager{},
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

// ListLifecycleWorkflowWorkflowRunUserProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowWorkflowRunUserProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId) (ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowRunUserProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowRunUserProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
