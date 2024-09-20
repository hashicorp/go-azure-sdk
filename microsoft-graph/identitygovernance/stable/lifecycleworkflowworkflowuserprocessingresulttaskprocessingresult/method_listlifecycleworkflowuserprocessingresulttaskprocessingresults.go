package lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult

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

type ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions() ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions {
	return ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions{}
}

func (o ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowUserProcessingResultTaskProcessingResults - List taskProcessingResults. Get the task processing
// result from a userProcessingResult either directly or through a run.
func (c LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient) ListLifecycleWorkflowUserProcessingResultTaskProcessingResults(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, options ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions) (result ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCustomPager{},
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
		Values *[]stable.IdentityGovernanceTaskProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient) ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, options ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions) (ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient) ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultId, options ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults(ctx, id, options)
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

	result = ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
