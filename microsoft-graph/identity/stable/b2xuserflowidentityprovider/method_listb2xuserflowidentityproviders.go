package b2xuserflowidentityprovider

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

type ListB2xUserFlowIdentityProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityProvider
}

type ListB2xUserFlowIdentityProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityProvider
}

type ListB2xUserFlowIdentityProvidersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListB2xUserFlowIdentityProvidersOperationOptions() ListB2xUserFlowIdentityProvidersOperationOptions {
	return ListB2xUserFlowIdentityProvidersOperationOptions{}
}

func (o ListB2xUserFlowIdentityProvidersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListB2xUserFlowIdentityProvidersOperationOptions) ToOData() *odata.Query {
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

func (o ListB2xUserFlowIdentityProvidersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListB2xUserFlowIdentityProvidersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowIdentityProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowIdentityProviders - List identityProviders. Get the identity providers in a b2xIdentityUserFlow
// object.
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProviders(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProvidersOperationOptions) (result ListB2xUserFlowIdentityProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListB2xUserFlowIdentityProvidersCustomPager{},
		Path:          fmt.Sprintf("%s/identityProviders", id.ID()),
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
		Values *[]stable.IdentityProvider `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowIdentityProvidersComplete retrieves all the results into a single object
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProvidersComplete(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProvidersOperationOptions) (ListB2xUserFlowIdentityProvidersCompleteResult, error) {
	return c.ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate(ctx, id, options, IdentityProviderOperationPredicate{})
}

// ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowIdentityProviderClient) ListB2xUserFlowIdentityProvidersCompleteMatchingPredicate(ctx context.Context, id stable.IdentityB2xUserFlowId, options ListB2xUserFlowIdentityProvidersOperationOptions, predicate IdentityProviderOperationPredicate) (result ListB2xUserFlowIdentityProvidersCompleteResult, err error) {
	items := make([]stable.IdentityProvider, 0)

	resp, err := c.ListB2xUserFlowIdentityProviders(ctx, id, options)
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

	result = ListB2xUserFlowIdentityProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
