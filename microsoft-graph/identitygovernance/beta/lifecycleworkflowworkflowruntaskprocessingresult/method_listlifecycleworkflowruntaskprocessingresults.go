package lifecycleworkflowworkflowruntaskprocessingresult

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

type ListLifecycleWorkflowRunTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowRunTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListLifecycleWorkflowRunTaskProcessingResultsOperationOptions() ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions {
	return ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions{}
}

func (o ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowRunTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowRunTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowRunTaskProcessingResults - List taskProcessingResults. Get the taskProcessingResult resources
// for a run.
func (c LifecycleWorkflowWorkflowRunTaskProcessingResultClient) ListLifecycleWorkflowRunTaskProcessingResults(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions) (result ListLifecycleWorkflowRunTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowRunTaskProcessingResultsCustomPager{},
		Path:          fmt.Sprintf("%s/taskProcessingResults", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListLifecycleWorkflowRunTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowRunTaskProcessingResultClient) ListLifecycleWorkflowRunTaskProcessingResultsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions) (ListLifecycleWorkflowRunTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowRunTaskProcessingResultsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowRunTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowRunTaskProcessingResultClient) ListLifecycleWorkflowRunTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunTaskProcessingResultsOperationOptions, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowRunTaskProcessingResultsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowRunTaskProcessingResults(ctx, id, options)
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

	result = ListLifecycleWorkflowRunTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
