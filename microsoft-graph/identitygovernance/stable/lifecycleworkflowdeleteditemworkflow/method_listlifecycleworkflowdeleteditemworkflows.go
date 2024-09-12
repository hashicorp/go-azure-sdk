package lifecycleworkflowdeleteditemworkflow

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

type ListLifecycleWorkflowDeletedItemWorkflowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowDeletedItemWorkflowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflow
}

type ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListLifecycleWorkflowDeletedItemWorkflowsOperationOptions() ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions {
	return ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions{}
}

func (o ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowDeletedItemWorkflowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflows - List deletedItems (deleted lifecycle workflows). Get a list of the
// deleted workflow objects and their properties.
func (c LifecycleWorkflowDeletedItemWorkflowClient) ListLifecycleWorkflowDeletedItemWorkflows(ctx context.Context, options ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions) (result ListLifecycleWorkflowDeletedItemWorkflowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowDeletedItemWorkflowsCustomPager{},
		Path:          "/identityGovernance/lifecycleWorkflows/deletedItems/workflows",
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
		Values *[]stable.IdentityGovernanceWorkflow `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowClient) ListLifecycleWorkflowDeletedItemWorkflowsComplete(ctx context.Context, options ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions) (ListLifecycleWorkflowDeletedItemWorkflowsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowsCompleteMatchingPredicate(ctx, options, IdentityGovernanceWorkflowOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowClient) ListLifecycleWorkflowDeletedItemWorkflowsCompleteMatchingPredicate(ctx context.Context, options ListLifecycleWorkflowDeletedItemWorkflowsOperationOptions, predicate IdentityGovernanceWorkflowOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflow, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflows(ctx, options)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
