package marketplaceregistrationdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WithoutScopeListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MarketplaceRegistrationDefinition
}

type WithoutScopeListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MarketplaceRegistrationDefinition
}

type WithoutScopeListOperationOptions struct {
	Filter *string
}

func DefaultWithoutScopeListOperationOptions() WithoutScopeListOperationOptions {
	return WithoutScopeListOperationOptions{}
}

func (o WithoutScopeListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WithoutScopeListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WithoutScopeListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type WithoutScopeListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WithoutScopeListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WithoutScopeList ...
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeList(ctx context.Context, options WithoutScopeListOperationOptions) (result WithoutScopeListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &WithoutScopeListCustomPager{},
		Path:          "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
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
		Values *[]MarketplaceRegistrationDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WithoutScopeListComplete retrieves all the results into a single object
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeListComplete(ctx context.Context, options WithoutScopeListOperationOptions) (WithoutScopeListCompleteResult, error) {
	return c.WithoutScopeListCompleteMatchingPredicate(ctx, options, MarketplaceRegistrationDefinitionOperationPredicate{})
}

// WithoutScopeListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeListCompleteMatchingPredicate(ctx context.Context, options WithoutScopeListOperationOptions, predicate MarketplaceRegistrationDefinitionOperationPredicate) (result WithoutScopeListCompleteResult, err error) {
	items := make([]MarketplaceRegistrationDefinition, 0)

	resp, err := c.WithoutScopeList(ctx, options)
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

	result = WithoutScopeListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
