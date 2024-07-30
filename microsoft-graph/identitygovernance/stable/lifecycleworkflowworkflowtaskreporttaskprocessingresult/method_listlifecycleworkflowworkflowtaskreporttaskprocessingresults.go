package lifecycleworkflowworkflowtaskreporttaskprocessingresult

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

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResults ...
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) (result ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCustomPager{},
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

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId) (ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportId, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
