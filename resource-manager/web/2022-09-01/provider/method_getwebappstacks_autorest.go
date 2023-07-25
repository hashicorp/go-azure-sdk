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

type GetWebAppStacksOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebAppStack

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetWebAppStacksOperationResponse, error)
}

type GetWebAppStacksCompleteResult struct {
	Items []WebAppStack
}

func (r GetWebAppStacksOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetWebAppStacksOperationResponse) LoadMore(ctx context.Context) (resp GetWebAppStacksOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetWebAppStacksOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetWebAppStacksOperationOptions() GetWebAppStacksOperationOptions {
	return GetWebAppStacksOperationOptions{}
}

func (o GetWebAppStacksOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetWebAppStacksOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.StackOsType != nil {
		out["stackOsType"] = *o.StackOsType
	}

	return out
}

// GetWebAppStacks ...
func (c ProviderClient) GetWebAppStacks(ctx context.Context, options GetWebAppStacksOperationOptions) (resp GetWebAppStacksOperationResponse, err error) {
	req, err := c.preparerForGetWebAppStacks(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetWebAppStacks(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetWebAppStacks prepares the GetWebAppStacks request.
func (c ProviderClient) preparerForGetWebAppStacks(ctx context.Context, options GetWebAppStacksOperationOptions) (*http.Request, error) {
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
		autorest.WithPath("/providers/Microsoft.Web/webAppStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetWebAppStacksWithNextLink prepares the GetWebAppStacks request with the given nextLink token.
func (c ProviderClient) preparerForGetWebAppStacksWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetWebAppStacks handles the response to the GetWebAppStacks request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetWebAppStacks(resp *http.Response) (result GetWebAppStacksOperationResponse, err error) {
	type page struct {
		Values   []WebAppStack `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetWebAppStacksOperationResponse, err error) {
			req, err := c.preparerForGetWebAppStacksWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetWebAppStacks(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacks", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetWebAppStacksComplete retrieves all of the results into a single object
func (c ProviderClient) GetWebAppStacksComplete(ctx context.Context, options GetWebAppStacksOperationOptions) (GetWebAppStacksCompleteResult, error) {
	return c.GetWebAppStacksCompleteMatchingPredicate(ctx, options, WebAppStackOperationPredicate{})
}

// GetWebAppStacksCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetWebAppStacksCompleteMatchingPredicate(ctx context.Context, options GetWebAppStacksOperationOptions, predicate WebAppStackOperationPredicate) (resp GetWebAppStacksCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	page, err := c.GetWebAppStacks(ctx, options)
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

	out := GetWebAppStacksCompleteResult{
		Items: items,
	}
	return out, nil
}
