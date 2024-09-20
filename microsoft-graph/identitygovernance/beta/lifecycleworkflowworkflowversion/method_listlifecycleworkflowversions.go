package lifecycleworkflowworkflowversion

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

type ListLifecycleWorkflowVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowVersionsOperationOptions struct {
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

func DefaultListLifecycleWorkflowVersionsOperationOptions() ListLifecycleWorkflowVersionsOperationOptions {
	return ListLifecycleWorkflowVersionsOperationOptions{}
}

func (o ListLifecycleWorkflowVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowVersions - List workflowVersions. Get a list of the workflowVersion objects and their
// properties.
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowVersions(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowVersionsOperationOptions) (result ListLifecycleWorkflowVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]beta.IdentityGovernanceWorkflowVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowVersionsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowVersionsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowVersionsOperationOptions) (ListLifecycleWorkflowVersionsCompleteResult, error) {
	return c.ListLifecycleWorkflowVersionsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceWorkflowVersionOperationPredicate{})
}

// ListLifecycleWorkflowVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowVersionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowVersionsOperationOptions, predicate IdentityGovernanceWorkflowVersionOperationPredicate) (result ListLifecycleWorkflowVersionsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceWorkflowVersion, 0)

	resp, err := c.ListLifecycleWorkflowVersions(ctx, id, options)
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

	result = ListLifecycleWorkflowVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
