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

type ProviderGetFunctionAppStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FunctionAppStack
}

type ProviderGetFunctionAppStacksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FunctionAppStack
}

type ProviderGetFunctionAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultProviderGetFunctionAppStacksOperationOptions() ProviderGetFunctionAppStacksOperationOptions {
	return ProviderGetFunctionAppStacksOperationOptions{}
}

func (o ProviderGetFunctionAppStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProviderGetFunctionAppStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ProviderGetFunctionAppStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

type ProviderGetFunctionAppStacksCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProviderGetFunctionAppStacksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProviderGetFunctionAppStacks ...
func (c OpenapisClient) ProviderGetFunctionAppStacks(ctx context.Context, options ProviderGetFunctionAppStacksOperationOptions) (result ProviderGetFunctionAppStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ProviderGetFunctionAppStacksCustomPager{},
		Path:          "/providers/Microsoft.Web/functionAppStacks",
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
		Values *[]FunctionAppStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProviderGetFunctionAppStacksComplete retrieves all the results into a single object
func (c OpenapisClient) ProviderGetFunctionAppStacksComplete(ctx context.Context, options ProviderGetFunctionAppStacksOperationOptions) (ProviderGetFunctionAppStacksCompleteResult, error) {
	return c.ProviderGetFunctionAppStacksCompleteMatchingPredicate(ctx, options, FunctionAppStackOperationPredicate{})
}

// ProviderGetFunctionAppStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ProviderGetFunctionAppStacksCompleteMatchingPredicate(ctx context.Context, options ProviderGetFunctionAppStacksOperationOptions, predicate FunctionAppStackOperationPredicate) (result ProviderGetFunctionAppStacksCompleteResult, err error) {
	items := make([]FunctionAppStack, 0)

	resp, err := c.ProviderGetFunctionAppStacks(ctx, options)
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

	result = ProviderGetFunctionAppStacksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
