package identityprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityProviderBase
}

type ListProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityProviderBase
}

type ListProvidersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListProvidersOperationOptions() ListProvidersOperationOptions {
	return ListProvidersOperationOptions{}
}

func (o ListProvidersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProvidersOperationOptions) ToOData() *odata.Query {
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

func (o ListProvidersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProvidersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProviders - List identityProviders. Get a collection of identity provider resources that are configured for a
// tenant, and that are derived from identityProviderBase. For a Microsoft Entra tenant, the providers can be
// socialIdentityProviders or builtinIdentityProviders objects. For an Azure AD B2C, the providers can be
// socialIdentityProvider, or appleManagedIdentityProvider objects.
func (c IdentityProviderClient) ListProviders(ctx context.Context, options ListProvidersOperationOptions) (result ListProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProvidersCustomPager{},
		Path:          "/identity/identityProviders",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.IdentityProviderBase, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalIdentityProviderBaseImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.IdentityProviderBase (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListProvidersComplete retrieves all the results into a single object
func (c IdentityProviderClient) ListProvidersComplete(ctx context.Context, options ListProvidersOperationOptions) (ListProvidersCompleteResult, error) {
	return c.ListProvidersCompleteMatchingPredicate(ctx, options, IdentityProviderBaseOperationPredicate{})
}

// ListProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityProviderClient) ListProvidersCompleteMatchingPredicate(ctx context.Context, options ListProvidersOperationOptions, predicate IdentityProviderBaseOperationPredicate) (result ListProvidersCompleteResult, err error) {
	items := make([]stable.IdentityProviderBase, 0)

	resp, err := c.ListProviders(ctx, options)
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

	result = ListProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
