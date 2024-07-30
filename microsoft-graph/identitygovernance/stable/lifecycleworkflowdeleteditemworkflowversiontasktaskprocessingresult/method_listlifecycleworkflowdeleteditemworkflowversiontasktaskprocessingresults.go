package lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult

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

type ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResults ...
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId) (result ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCustomPager{},
		Path:       fmt.Sprintf("%s/taskProcessingResults", id.ID()),
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
		Values *[]stable.IdentityGovernanceTaskProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId) (ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
