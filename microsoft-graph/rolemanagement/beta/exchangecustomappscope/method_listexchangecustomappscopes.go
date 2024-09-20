package exchangecustomappscope

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

type ListExchangeCustomAppScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CustomAppScope
}

type ListExchangeCustomAppScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CustomAppScope
}

type ListExchangeCustomAppScopesOperationOptions struct {
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

func DefaultListExchangeCustomAppScopesOperationOptions() ListExchangeCustomAppScopesOperationOptions {
	return ListExchangeCustomAppScopesOperationOptions{}
}

func (o ListExchangeCustomAppScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeCustomAppScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListExchangeCustomAppScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeCustomAppScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeCustomAppScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeCustomAppScopes - List customAppScopes. Get a list of customAppScope objects for an RBAC provider.
// Currently only the Exchange Online RBAC provider is supported.
func (c ExchangeCustomAppScopeClient) ListExchangeCustomAppScopes(ctx context.Context, options ListExchangeCustomAppScopesOperationOptions) (result ListExchangeCustomAppScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeCustomAppScopesCustomPager{},
		Path:          "/roleManagement/exchange/customAppScopes",
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
		Values *[]beta.CustomAppScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeCustomAppScopesComplete retrieves all the results into a single object
func (c ExchangeCustomAppScopeClient) ListExchangeCustomAppScopesComplete(ctx context.Context, options ListExchangeCustomAppScopesOperationOptions) (ListExchangeCustomAppScopesCompleteResult, error) {
	return c.ListExchangeCustomAppScopesCompleteMatchingPredicate(ctx, options, CustomAppScopeOperationPredicate{})
}

// ListExchangeCustomAppScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeCustomAppScopeClient) ListExchangeCustomAppScopesCompleteMatchingPredicate(ctx context.Context, options ListExchangeCustomAppScopesOperationOptions, predicate CustomAppScopeOperationPredicate) (result ListExchangeCustomAppScopesCompleteResult, err error) {
	items := make([]beta.CustomAppScope, 0)

	resp, err := c.ListExchangeCustomAppScopes(ctx, options)
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

	result = ListExchangeCustomAppScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
