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

type GetWebAppStacksForLocationOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebAppStack

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetWebAppStacksForLocationOperationResponse, error)
}

type GetWebAppStacksForLocationCompleteResult struct {
	Items []WebAppStack
}

func (r GetWebAppStacksForLocationOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetWebAppStacksForLocationOperationResponse) LoadMore(ctx context.Context) (resp GetWebAppStacksForLocationOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetWebAppStacksForLocationOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetWebAppStacksForLocationOperationOptions() GetWebAppStacksForLocationOperationOptions {
	return GetWebAppStacksForLocationOperationOptions{}
}

func (o GetWebAppStacksForLocationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetWebAppStacksForLocationOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.StackOsType != nil {
		out["stackOsType"] = *o.StackOsType
	}

	return out
}

// GetWebAppStacksForLocation ...
func (c ProviderClient) GetWebAppStacksForLocation(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions) (resp GetWebAppStacksForLocationOperationResponse, err error) {
	req, err := c.preparerForGetWebAppStacksForLocation(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetWebAppStacksForLocation(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetWebAppStacksForLocation prepares the GetWebAppStacksForLocation request.
func (c ProviderClient) preparerForGetWebAppStacksForLocation(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/webAppStacks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetWebAppStacksForLocationWithNextLink prepares the GetWebAppStacksForLocation request with the given nextLink token.
func (c ProviderClient) preparerForGetWebAppStacksForLocationWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetWebAppStacksForLocation handles the response to the GetWebAppStacksForLocation request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetWebAppStacksForLocation(resp *http.Response) (result GetWebAppStacksForLocationOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetWebAppStacksForLocationOperationResponse, err error) {
			req, err := c.preparerForGetWebAppStacksForLocationWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetWebAppStacksForLocation(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetWebAppStacksForLocation", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetWebAppStacksForLocationComplete retrieves all of the results into a single object
func (c ProviderClient) GetWebAppStacksForLocationComplete(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions) (GetWebAppStacksForLocationCompleteResult, error) {
	return c.GetWebAppStacksForLocationCompleteMatchingPredicate(ctx, id, options, WebAppStackOperationPredicate{})
}

// GetWebAppStacksForLocationCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetWebAppStacksForLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions, predicate WebAppStackOperationPredicate) (resp GetWebAppStacksForLocationCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	page, err := c.GetWebAppStacksForLocation(ctx, id, options)
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

	out := GetWebAppStacksForLocationCompleteResult{
		Items: items,
	}
	return out, nil
}
