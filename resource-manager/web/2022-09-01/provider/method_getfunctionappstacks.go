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

type GetFunctionAppStacksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FunctionAppStack
}

type GetFunctionAppStacksCompleteResult struct {
	Items []FunctionAppStack
}

type GetFunctionAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetFunctionAppStacksOperationOptions() GetFunctionAppStacksOperationOptions {
	return GetFunctionAppStacksOperationOptions{}
}

func (o GetFunctionAppStacksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetFunctionAppStacksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetFunctionAppStacksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

// GetFunctionAppStacks ...
func (c ProviderClient) GetFunctionAppStacks(ctx context.Context, options GetFunctionAppStacksOperationOptions) (result GetFunctionAppStacksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          "/providers/Microsoft.Web/functionAppStacks",
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
		Values *[]FunctionAppStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetFunctionAppStacksComplete retrieves all the results into a single object
func (c ProviderClient) GetFunctionAppStacksComplete(ctx context.Context, options GetFunctionAppStacksOperationOptions) (GetFunctionAppStacksCompleteResult, error) {
	return c.GetFunctionAppStacksCompleteMatchingPredicate(ctx, options, FunctionAppStackOperationPredicate{})
}

// GetFunctionAppStacksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetFunctionAppStacksCompleteMatchingPredicate(ctx context.Context, options GetFunctionAppStacksOperationOptions, predicate FunctionAppStackOperationPredicate) (result GetFunctionAppStacksCompleteResult, err error) {
	items := make([]FunctionAppStack, 0)

	resp, err := c.GetFunctionAppStacks(ctx, options)
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

	result = GetFunctionAppStacksCompleteResult{
		Items: items,
	}
	return
}
