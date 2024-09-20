package lifecycleworkflowworkflowtaskreport

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

type ListLifecycleWorkflowTaskReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskReport
}

type ListLifecycleWorkflowTaskReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskReport
}

type ListLifecycleWorkflowTaskReportsOperationOptions struct {
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

func DefaultListLifecycleWorkflowTaskReportsOperationOptions() ListLifecycleWorkflowTaskReportsOperationOptions {
	return ListLifecycleWorkflowTaskReportsOperationOptions{}
}

func (o ListLifecycleWorkflowTaskReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowTaskReportsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowTaskReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowTaskReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTaskReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTaskReports - List taskReports. Get a list of the taskReport objects and their properties.
func (c LifecycleWorkflowWorkflowTaskReportClient) ListLifecycleWorkflowTaskReports(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowTaskReportsOperationOptions) (result ListLifecycleWorkflowTaskReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowTaskReportsCustomPager{},
		Path:          fmt.Sprintf("%s/taskReports", id.ID()),
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
		Values *[]stable.IdentityGovernanceTaskReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowTaskReportsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskReportClient) ListLifecycleWorkflowTaskReportsComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowTaskReportsOperationOptions) (ListLifecycleWorkflowTaskReportsCompleteResult, error) {
	return c.ListLifecycleWorkflowTaskReportsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskReportOperationPredicate{})
}

// ListLifecycleWorkflowTaskReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskReportClient) ListLifecycleWorkflowTaskReportsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowTaskReportsOperationOptions, predicate IdentityGovernanceTaskReportOperationPredicate) (result ListLifecycleWorkflowTaskReportsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskReport, 0)

	resp, err := c.ListLifecycleWorkflowTaskReports(ctx, id, options)
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

	result = ListLifecycleWorkflowTaskReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
