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

type ProviderGetWebAppStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebAppStack
}

type ProviderGetWebAppStacksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebAppStack
}

type ProviderGetWebAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultProviderGetWebAppStacksOperationOptions() ProviderGetWebAppStacksOperationOptions {
	return ProviderGetWebAppStacksOperationOptions{}
}

func (o ProviderGetWebAppStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProviderGetWebAppStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ProviderGetWebAppStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

type ProviderGetWebAppStacksCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProviderGetWebAppStacksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProviderGetWebAppStacks ...
func (c OpenapisClient) ProviderGetWebAppStacks(ctx context.Context, options ProviderGetWebAppStacksOperationOptions) (result ProviderGetWebAppStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ProviderGetWebAppStacksCustomPager{},
		Path:          "/providers/Microsoft.Web/webAppStacks",
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
		Values *[]WebAppStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ProviderGetWebAppStacksComplete retrieves all the results into a single object
func (c OpenapisClient) ProviderGetWebAppStacksComplete(ctx context.Context, options ProviderGetWebAppStacksOperationOptions) (ProviderGetWebAppStacksCompleteResult, error) {
	return c.ProviderGetWebAppStacksCompleteMatchingPredicate(ctx, options, WebAppStackOperationPredicate{})
}

// ProviderGetWebAppStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ProviderGetWebAppStacksCompleteMatchingPredicate(ctx context.Context, options ProviderGetWebAppStacksOperationOptions, predicate WebAppStackOperationPredicate) (result ProviderGetWebAppStacksCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	resp, err := c.ProviderGetWebAppStacks(ctx, options)
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

	result = ProviderGetWebAppStacksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
