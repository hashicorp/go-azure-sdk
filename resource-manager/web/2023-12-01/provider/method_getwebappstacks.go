package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWebAppStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebAppStack
}

type GetWebAppStacksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebAppStack
}

type GetWebAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetWebAppStacksOperationOptions() GetWebAppStacksOperationOptions {
	return GetWebAppStacksOperationOptions{}
}

func (o GetWebAppStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetWebAppStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetWebAppStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

// GetWebAppStacks ...
func (c ProviderClient) GetWebAppStacks(ctx context.Context, options GetWebAppStacksOperationOptions) (result GetWebAppStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          "/providers/Microsoft.Web/webAppStacks",
		OptionsObject: options,
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

// GetWebAppStacksComplete retrieves all the results into a single object
func (c ProviderClient) GetWebAppStacksComplete(ctx context.Context, options GetWebAppStacksOperationOptions) (GetWebAppStacksCompleteResult, error) {
	return c.GetWebAppStacksCompleteMatchingPredicate(ctx, options, WebAppStackOperationPredicate{})
}

// GetWebAppStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetWebAppStacksCompleteMatchingPredicate(ctx context.Context, options GetWebAppStacksOperationOptions, predicate WebAppStackOperationPredicate) (result GetWebAppStacksCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	resp, err := c.GetWebAppStacks(ctx, options)
	if err != nil {
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

	result = GetWebAppStacksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
