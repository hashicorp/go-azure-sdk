package lifecycleworkflowworkflowtemplatetasktaskprocessingresult

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

type ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResults ...
func (c LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResults(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) (result ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCustomPager{},
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

// ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId) (ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient) ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResults(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
