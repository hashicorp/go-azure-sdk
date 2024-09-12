package lifecycleworkflowrun

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

type ListLifecycleWorkflowRunsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceRun
}

type ListLifecycleWorkflowRunsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceRun
}

type ListLifecycleWorkflowRunsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListLifecycleWorkflowRunsOperationOptions() ListLifecycleWorkflowRunsOperationOptions {
	return ListLifecycleWorkflowRunsOperationOptions{}
}

func (o ListLifecycleWorkflowRunsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowRunsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowRunsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowRunsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowRunsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowRuns - List runs. Get a list of the run objects and their properties for a lifecycle workflow.
func (c LifecycleWorkflowRunClient) ListLifecycleWorkflowRuns(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowRunsOperationOptions) (result ListLifecycleWorkflowRunsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowRunsCustomPager{},
		Path:          fmt.Sprintf("%s/runs", id.ID()),
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
		Values *[]beta.IdentityGovernanceRun `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowRunsComplete retrieves all the results into a single object
func (c LifecycleWorkflowRunClient) ListLifecycleWorkflowRunsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowRunsOperationOptions) (ListLifecycleWorkflowRunsCompleteResult, error) {
	return c.ListLifecycleWorkflowRunsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceRunOperationPredicate{})
}

// ListLifecycleWorkflowRunsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowRunClient) ListLifecycleWorkflowRunsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowRunsOperationOptions, predicate IdentityGovernanceRunOperationPredicate) (result ListLifecycleWorkflowRunsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceRun, 0)

	resp, err := c.ListLifecycleWorkflowRuns(ctx, id, options)
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

	result = ListLifecycleWorkflowRunsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
