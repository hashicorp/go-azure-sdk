package lifecycleworkflowworkflowtasktaskprocessingresult

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

type ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTaskTaskProcessingResults ...
func (c LifecycleWorkflowWorkflowTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskTaskProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) (result ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCustomPager{},
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

// ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId) (ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTaskTaskProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowTaskTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
