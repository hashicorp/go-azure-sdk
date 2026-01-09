package lifecycleworkflowworkflowtemplate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLifecycleWorkflowTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflowTemplate
}

type ListLifecycleWorkflowTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflowTemplate
}

type ListLifecycleWorkflowTemplatesOperationOptions struct {
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

func DefaultListLifecycleWorkflowTemplatesOperationOptions() ListLifecycleWorkflowTemplatesOperationOptions {
	return ListLifecycleWorkflowTemplatesOperationOptions{}
}

func (o ListLifecycleWorkflowTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTemplates - List workflowTemplates. Get a list of the workflowTemplate objects and their
// properties.
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowTemplates(ctx context.Context, options ListLifecycleWorkflowTemplatesOperationOptions) (result ListLifecycleWorkflowTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowTemplatesCustomPager{},
		Path:          "/identityGovernance/lifecycleWorkflows/workflowTemplates",
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
		Values *[]stable.IdentityGovernanceWorkflowTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowTemplatesComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowTemplatesComplete(ctx context.Context, options ListLifecycleWorkflowTemplatesOperationOptions) (ListLifecycleWorkflowTemplatesCompleteResult, error) {
	return c.ListLifecycleWorkflowTemplatesCompleteMatchingPredicate(ctx, options, IdentityGovernanceWorkflowTemplateOperationPredicate{})
}

// ListLifecycleWorkflowTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowTemplatesCompleteMatchingPredicate(ctx context.Context, options ListLifecycleWorkflowTemplatesOperationOptions, predicate IdentityGovernanceWorkflowTemplateOperationPredicate) (result ListLifecycleWorkflowTemplatesCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflowTemplate, 0)

	resp, err := c.ListLifecycleWorkflowTemplates(ctx, options)
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

	result = ListLifecycleWorkflowTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
