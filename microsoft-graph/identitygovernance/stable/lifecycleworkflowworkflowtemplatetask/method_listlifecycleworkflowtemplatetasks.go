package lifecycleworkflowworkflowtemplatetask

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

type ListLifecycleWorkflowTemplateTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowTemplateTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTask
}

type ListLifecycleWorkflowTemplateTasksOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListLifecycleWorkflowTemplateTasksOperationOptions() ListLifecycleWorkflowTemplateTasksOperationOptions {
	return ListLifecycleWorkflowTemplateTasksOperationOptions{}
}

func (o ListLifecycleWorkflowTemplateTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowTemplateTasksOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowTemplateTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowTemplateTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTemplateTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTemplateTasks - Get tasks from identityGovernance. Represents the configured tasks to execute
// and their execution sequence within a workflow. This relationship is expanded by default.
func (c LifecycleWorkflowWorkflowTemplateTaskClient) ListLifecycleWorkflowTemplateTasks(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowTemplateId, options ListLifecycleWorkflowTemplateTasksOperationOptions) (result ListLifecycleWorkflowTemplateTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowTemplateTasksCustomPager{},
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
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
		Values *[]stable.IdentityGovernanceTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowTemplateTasksComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTemplateTaskClient) ListLifecycleWorkflowTemplateTasksComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowTemplateId, options ListLifecycleWorkflowTemplateTasksOperationOptions) (ListLifecycleWorkflowTemplateTasksCompleteResult, error) {
	return c.ListLifecycleWorkflowTemplateTasksCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskOperationPredicate{})
}

// ListLifecycleWorkflowTemplateTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTemplateTaskClient) ListLifecycleWorkflowTemplateTasksCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowTemplateId, options ListLifecycleWorkflowTemplateTasksOperationOptions, predicate IdentityGovernanceTaskOperationPredicate) (result ListLifecycleWorkflowTemplateTasksCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTask, 0)

	resp, err := c.ListLifecycleWorkflowTemplateTasks(ctx, id, options)
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

	result = ListLifecycleWorkflowTemplateTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
