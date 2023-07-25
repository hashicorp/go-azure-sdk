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

type GetFunctionAppStacksForLocationOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FunctionAppStack

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetFunctionAppStacksForLocationOperationResponse, error)
}

type GetFunctionAppStacksForLocationCompleteResult struct {
	Items []FunctionAppStack
}

func (r GetFunctionAppStacksForLocationOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetFunctionAppStacksForLocationOperationResponse) LoadMore(ctx context.Context) (resp GetFunctionAppStacksForLocationOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetFunctionAppStacksForLocationOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetFunctionAppStacksForLocationOperationOptions() GetFunctionAppStacksForLocationOperationOptions {
	return GetFunctionAppStacksForLocationOperationOptions{}
}

func (o GetFunctionAppStacksForLocationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetFunctionAppStacksForLocationOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.StackOsType != nil {
		out["stackOsType"] = *o.StackOsType
	}

	return out
}

// GetFunctionAppStacksForLocation ...
func (c ProviderClient) GetFunctionAppStacksForLocation(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions) (resp GetFunctionAppStacksForLocationOperationResponse, err error) {
	req, err := c.preparerForGetFunctionAppStacksForLocation(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetFunctionAppStacksForLocation(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetFunctionAppStacksForLocation prepares the GetFunctionAppStacksForLocation request.
func (c ProviderClient) preparerForGetFunctionAppStacksForLocation(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/functionAppStacks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetFunctionAppStacksForLocationWithNextLink prepares the GetFunctionAppStacksForLocation request with the given nextLink token.
func (c ProviderClient) preparerForGetFunctionAppStacksForLocationWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetFunctionAppStacksForLocation handles the response to the GetFunctionAppStacksForLocation request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetFunctionAppStacksForLocation(resp *http.Response) (result GetFunctionAppStacksForLocationOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetFunctionAppStacksForLocationOperationResponse, err error) {
			req, err := c.preparerForGetFunctionAppStacksForLocationWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetFunctionAppStacksForLocation(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetFunctionAppStacksForLocation", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetFunctionAppStacksForLocationComplete retrieves all of the results into a single object
func (c ProviderClient) GetFunctionAppStacksForLocationComplete(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions) (GetFunctionAppStacksForLocationCompleteResult, error) {
	return c.GetFunctionAppStacksForLocationCompleteMatchingPredicate(ctx, id, options, FunctionAppStackOperationPredicate{})
}

// GetFunctionAppStacksForLocationCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetFunctionAppStacksForLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions, predicate FunctionAppStackOperationPredicate) (resp GetFunctionAppStacksForLocationCompleteResult, err error) {
	items := make([]FunctionAppStack, 0)

	page, err := c.GetFunctionAppStacksForLocation(ctx, id, options)
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

	out := GetFunctionAppStacksForLocationCompleteResult{
		Items: items,
	}
	return out, nil
}
