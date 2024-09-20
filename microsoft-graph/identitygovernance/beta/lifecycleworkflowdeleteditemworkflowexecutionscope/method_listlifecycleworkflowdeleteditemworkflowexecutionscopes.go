package lifecycleworkflowdeleteditemworkflowexecutionscope

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

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions struct {
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

func DefaultListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions() ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions {
	return ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions{}
}

func (o ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes - Get executionScope from identityGovernance. The unique
// identifier of the Microsoft Entra identity that last modified the workflow object.
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions) (result ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCustomPager{},
		Path:          fmt.Sprintf("%s/executionScope", id.ID()),
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
		Values *[]beta.IdentityGovernanceUserProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions) (ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient) ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, options ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesOperationOptions, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowExecutionScopes(ctx, id, options)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowExecutionScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
