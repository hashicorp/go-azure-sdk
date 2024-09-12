package lifecycleworkflowtaskprocessingresult

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLifecycleWorkflowTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowTaskProcessingResultsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListLifecycleWorkflowTaskProcessingResultsOperationOptions() ListLifecycleWorkflowTaskProcessingResultsOperationOptions {
	return ListLifecycleWorkflowTaskProcessingResultsOperationOptions{}
}

func (o ListLifecycleWorkflowTaskProcessingResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowTaskProcessingResultsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListLifecycleWorkflowTaskProcessingResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTaskProcessingResults - Get taskProcessingResults from identityGovernance. The result of
// processing the task.
func (c LifecycleWorkflowTaskProcessingResultClient) ListLifecycleWorkflowTaskProcessingResults(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, options ListLifecycleWorkflowTaskProcessingResultsOperationOptions) (result ListLifecycleWorkflowTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowTaskProcessingResultsCustomPager{},
		Path:          fmt.Sprintf("%s/taskProcessingResults", id.ID()),
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
		Values *[]beta.IdentityGovernanceTaskProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowTaskProcessingResultClient) ListLifecycleWorkflowTaskProcessingResultsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, options ListLifecycleWorkflowTaskProcessingResultsOperationOptions) (ListLifecycleWorkflowTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowTaskProcessingResultsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowTaskProcessingResultClient) ListLifecycleWorkflowTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskId, options ListLifecycleWorkflowTaskProcessingResultsOperationOptions, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowTaskProcessingResultsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowTaskProcessingResults(ctx, id, options)
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

	result = ListLifecycleWorkflowTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
