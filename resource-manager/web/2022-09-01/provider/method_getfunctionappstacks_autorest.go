package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetFunctionAppStacksOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FunctionAppStack

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetFunctionAppStacksOperationResponse, error)
}

type GetFunctionAppStacksCompleteResult struct {
	Items []FunctionAppStack
}

func (r GetFunctionAppStacksOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetFunctionAppStacksOperationResponse) LoadMore(ctx context.Context) (resp GetFunctionAppStacksOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetFunctionAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetFunctionAppStacksOperationOptions() GetFunctionAppStacksOperationOptions {
	return GetFunctionAppStacksOperationOptions{}
}

func (o GetFunctionAppStacksOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetFunctionAppStacksOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.StackOsType != nil {
		out["stackOsType"] = *o.StackOsType
	}

	return out
}

// GetFunctionAppStacks ...
func (c ProviderClient) GetFunctionAppStacks(ctx context.Context, options GetFunctionAppStacksOperationOptions) (resp GetFunctionAppStacksOperationResponse, err error) {
	req, err := c.preparerForGetFunctionAppStacks(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetFunctionAppStacks(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetFunctionAppStacks prepares the GetFunctionAppStacks request.
func (c ProviderClient) preparerForGetFunctionAppStacks(ctx context.Context, options GetFunctionAppStacksOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath("/providers/Microsoft.Web/functionAppStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetFunctionAppStacksWithNextLink prepares the GetFunctionAppStacks request with the given nextLink token.
func (c ProviderClient) preparerForGetFunctionAppStacksWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFunctionAppStacks handles the response to the GetFunctionAppStacks request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetFunctionAppStacks(resp *http.Response) (result GetFunctionAppStacksOperationResponse, err error) {
	type page struct {
		Values   []FunctionAppStack `json:"value"`
		NextLink *string            `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetFunctionAppStacksOperationResponse, err error) {
			req, err := c.preparerForGetFunctionAppStacksWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetFunctionAppStacks(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacks", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetFunctionAppStacksComplete retrieves all of the results into a single object
func (c ProviderClient) GetFunctionAppStacksComplete(ctx context.Context, options GetFunctionAppStacksOperationOptions) (GetFunctionAppStacksCompleteResult, error) {
	return c.GetFunctionAppStacksCompleteMatchingPredicate(ctx, options, FunctionAppStackOperationPredicate{})
}

// GetFunctionAppStacksCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetFunctionAppStacksCompleteMatchingPredicate(ctx context.Context, options GetFunctionAppStacksOperationOptions, predicate FunctionAppStackOperationPredicate) (resp GetFunctionAppStacksCompleteResult, err error) {
	items := make([]FunctionAppStack, 0)

	page, err := c.GetFunctionAppStacks(ctx, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := GetFunctionAppStacksCompleteResult{
		Items: items,
	}
	return out, nil
}
