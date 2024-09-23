package lifecycleworkflowworkflow

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

type ListLifecycleWorkflowWorkflowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowWorkflowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowWorkflowsOperationOptions struct {
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

func DefaultListLifecycleWorkflowWorkflowsOperationOptions() ListLifecycleWorkflowWorkflowsOperationOptions {
	return ListLifecycleWorkflowWorkflowsOperationOptions{}
}

func (o ListLifecycleWorkflowWorkflowsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowWorkflowsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowWorkflowsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowWorkflowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflows - List workflows. Get a list of workflow resources that are associated with lifecycle
// workflows.
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflows(ctx context.Context, options ListLifecycleWorkflowWorkflowsOperationOptions) (result ListLifecycleWorkflowWorkflowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowWorkflowsCustomPager{},
		Path:          "/identityGovernance/lifecycleWorkflows/workflows",
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
		Values *[]beta.IdentityGovernanceWorkflow `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowWorkflowsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflowsComplete(ctx context.Context, options ListLifecycleWorkflowWorkflowsOperationOptions) (ListLifecycleWorkflowWorkflowsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate(ctx, options, IdentityGovernanceWorkflowOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowClient) ListLifecycleWorkflowWorkflowsCompleteMatchingPredicate(ctx context.Context, options ListLifecycleWorkflowWorkflowsOperationOptions, predicate IdentityGovernanceWorkflowOperationPredicate) (result ListLifecycleWorkflowWorkflowsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceWorkflow, 0)

	resp, err := c.ListLifecycleWorkflowWorkflows(ctx, options)
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

	result = ListLifecycleWorkflowWorkflowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
