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

type GetAvailableStacksOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApplicationStackResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAvailableStacksOperationResponse, error)
}

type GetAvailableStacksCompleteResult struct {
	Items []ApplicationStackResource
}

func (r GetAvailableStacksOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAvailableStacksOperationResponse) LoadMore(ctx context.Context) (resp GetAvailableStacksOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetAvailableStacksOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultGetAvailableStacksOperationOptions() GetAvailableStacksOperationOptions {
	return GetAvailableStacksOperationOptions{}
}

func (o GetAvailableStacksOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAvailableStacksOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.OsTypeSelected != nil {
		out["osTypeSelected"] = *o.OsTypeSelected
	}

	return out
}

// GetAvailableStacks ...
func (c ProviderClient) GetAvailableStacks(ctx context.Context, options GetAvailableStacksOperationOptions) (resp GetAvailableStacksOperationResponse, err error) {
	req, err := c.preparerForGetAvailableStacks(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAvailableStacks(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAvailableStacks prepares the GetAvailableStacks request.
func (c ProviderClient) preparerForGetAvailableStacks(ctx context.Context, options GetAvailableStacksOperationOptions) (*http.Request, error) {
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
		autorest.WithPath("/providers/Microsoft.Web/availableStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAvailableStacksWithNextLink prepares the GetAvailableStacks request with the given nextLink token.
func (c ProviderClient) preparerForGetAvailableStacksWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAvailableStacks handles the response to the GetAvailableStacks request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetAvailableStacks(resp *http.Response) (result GetAvailableStacksOperationResponse, err error) {
	type page struct {
		Values   []ApplicationStackResource `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAvailableStacksOperationResponse, err error) {
			req, err := c.preparerForGetAvailableStacksWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAvailableStacks(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacks", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAvailableStacksComplete retrieves all of the results into a single object
func (c ProviderClient) GetAvailableStacksComplete(ctx context.Context, options GetAvailableStacksOperationOptions) (GetAvailableStacksCompleteResult, error) {
	return c.GetAvailableStacksCompleteMatchingPredicate(ctx, options, ApplicationStackResourceOperationPredicate{})
}

// GetAvailableStacksCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetAvailableStacksCompleteMatchingPredicate(ctx context.Context, options GetAvailableStacksOperationOptions, predicate ApplicationStackResourceOperationPredicate) (resp GetAvailableStacksCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	page, err := c.GetAvailableStacks(ctx, options)
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

	out := GetAvailableStacksCompleteResult{
		Items: items,
	}
	return out, nil
}
