package exchangeresourcenamespace

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

type ListExchangeResourceNamespacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceNamespace
}

type ListExchangeResourceNamespacesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceNamespace
}

type ListExchangeResourceNamespacesOperationOptions struct {
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

func DefaultListExchangeResourceNamespacesOperationOptions() ListExchangeResourceNamespacesOperationOptions {
	return ListExchangeResourceNamespacesOperationOptions{}
}

func (o ListExchangeResourceNamespacesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeResourceNamespacesOperationOptions) ToOData() *odata.Query {
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

func (o ListExchangeResourceNamespacesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeResourceNamespacesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeResourceNamespacesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeResourceNamespaces - Get resourceNamespaces from roleManagement. Resource that represents a collection of
// related actions.
func (c ExchangeResourceNamespaceClient) ListExchangeResourceNamespaces(ctx context.Context, options ListExchangeResourceNamespacesOperationOptions) (result ListExchangeResourceNamespacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeResourceNamespacesCustomPager{},
		Path:          "/roleManagement/exchange/resourceNamespaces",
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
		Values *[]beta.UnifiedRbacResourceNamespace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeResourceNamespacesComplete retrieves all the results into a single object
func (c ExchangeResourceNamespaceClient) ListExchangeResourceNamespacesComplete(ctx context.Context, options ListExchangeResourceNamespacesOperationOptions) (ListExchangeResourceNamespacesCompleteResult, error) {
	return c.ListExchangeResourceNamespacesCompleteMatchingPredicate(ctx, options, UnifiedRbacResourceNamespaceOperationPredicate{})
}

// ListExchangeResourceNamespacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeResourceNamespaceClient) ListExchangeResourceNamespacesCompleteMatchingPredicate(ctx context.Context, options ListExchangeResourceNamespacesOperationOptions, predicate UnifiedRbacResourceNamespaceOperationPredicate) (result ListExchangeResourceNamespacesCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceNamespace, 0)

	resp, err := c.ListExchangeResourceNamespaces(ctx, options)
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

	result = ListExchangeResourceNamespacesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
