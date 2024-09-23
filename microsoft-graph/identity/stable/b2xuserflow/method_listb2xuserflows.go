package b2xuserflow

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

type ListB2xUserFlowsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.B2xIdentityUserFlow
}

type ListB2xUserFlowsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.B2xIdentityUserFlow
}

type ListB2xUserFlowsOperationOptions struct {
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

func DefaultListB2xUserFlowsOperationOptions() ListB2xUserFlowsOperationOptions {
	return ListB2xUserFlowsOperationOptions{}
}

func (o ListB2xUserFlowsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowsOperationOptions) ToOData() *odata.Query {
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

func (o ListB2xUserFlowsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlows - List b2xIdentityUserFlows. Retrieve a list of b2xIdentityUserFlow objects.
func (c B2xUserFlowClient) ListB2xUserFlows(ctx context.Context, options ListB2xUserFlowsOperationOptions) (result ListB2xUserFlowsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowsCustomPager{},
		Path:          "/identity/b2xUserFlows",
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
		Values *[]stable.B2xIdentityUserFlow `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowsComplete retrieves all the results into a single object
func (c B2xUserFlowClient) ListB2xUserFlowsComplete(ctx context.Context, options ListB2xUserFlowsOperationOptions) (ListB2xUserFlowsCompleteResult, error) {
	return c.ListB2xUserFlowsCompleteMatchingPredicate(ctx, options, B2xIdentityUserFlowOperationPredicate{})
}

// ListB2xUserFlowsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowClient) ListB2xUserFlowsCompleteMatchingPredicate(ctx context.Context, options ListB2xUserFlowsOperationOptions, predicate B2xIdentityUserFlowOperationPredicate) (result ListB2xUserFlowsCompleteResult, err error) {
	items := make([]stable.B2xIdentityUserFlow, 0)

	resp, err := c.ListB2xUserFlows(ctx, options)
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

	result = ListB2xUserFlowsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
