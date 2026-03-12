package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderGetAvailableStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationStackResource
}

type ProviderGetAvailableStacksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationStackResource
}

type ProviderGetAvailableStacksOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultProviderGetAvailableStacksOperationOptions() ProviderGetAvailableStacksOperationOptions {
	return ProviderGetAvailableStacksOperationOptions{}
}

func (o ProviderGetAvailableStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProviderGetAvailableStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ProviderGetAvailableStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.OsTypeSelected != nil {
		out.Append("osTypeSelected", fmt.Sprintf("%v", *o.OsTypeSelected))
	}
	return &out
}

type ProviderGetAvailableStacksCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProviderGetAvailableStacksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProviderGetAvailableStacks ...
func (c OpenapisClient) ProviderGetAvailableStacks(ctx context.Context, options ProviderGetAvailableStacksOperationOptions) (result ProviderGetAvailableStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ProviderGetAvailableStacksCustomPager{},
		Path:          "/providers/Microsoft.Web/availableStacks",
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
		Values *[]ApplicationStackResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProviderGetAvailableStacksComplete retrieves all the results into a single object
func (c OpenapisClient) ProviderGetAvailableStacksComplete(ctx context.Context, options ProviderGetAvailableStacksOperationOptions) (ProviderGetAvailableStacksCompleteResult, error) {
	return c.ProviderGetAvailableStacksCompleteMatchingPredicate(ctx, options, ApplicationStackResourceOperationPredicate{})
}

// ProviderGetAvailableStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ProviderGetAvailableStacksCompleteMatchingPredicate(ctx context.Context, options ProviderGetAvailableStacksOperationOptions, predicate ApplicationStackResourceOperationPredicate) (result ProviderGetAvailableStacksCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	resp, err := c.ProviderGetAvailableStacks(ctx, options)
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

	result = ProviderGetAvailableStacksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
