package codeversion

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

type RegistryCodeVersionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CodeVersionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryCodeVersionsListOperationResponse, error)
}

type RegistryCodeVersionsListCompleteResult struct {
	Items []CodeVersionResource
}

func (r RegistryCodeVersionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryCodeVersionsListOperationResponse) LoadMore(ctx context.Context) (resp RegistryCodeVersionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryCodeVersionsListOperationOptions struct {
	OrderBy *string
	Skip    *string
	Top     *int64
}

func DefaultRegistryCodeVersionsListOperationOptions() RegistryCodeVersionsListOperationOptions {
	return RegistryCodeVersionsListOperationOptions{}
}

func (o RegistryCodeVersionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryCodeVersionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// RegistryCodeVersionsList ...
func (c CodeVersionClient) RegistryCodeVersionsList(ctx context.Context, id RegistryCodeId, options RegistryCodeVersionsListOperationOptions) (resp RegistryCodeVersionsListOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeVersionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryCodeVersionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryCodeVersionsList prepares the RegistryCodeVersionsList request.
func (c CodeVersionClient) preparerForRegistryCodeVersionsList(ctx context.Context, id RegistryCodeId, options RegistryCodeVersionsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryCodeVersionsListWithNextLink prepares the RegistryCodeVersionsList request with the given nextLink token.
func (c CodeVersionClient) preparerForRegistryCodeVersionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryCodeVersionsList handles the response to the RegistryCodeVersionsList request. The method always
// closes the http.Response Body.
func (c CodeVersionClient) responderForRegistryCodeVersionsList(resp *http.Response) (result RegistryCodeVersionsListOperationResponse, err error) {
	type page struct {
		Values   []CodeVersionResource `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryCodeVersionsListOperationResponse, err error) {
			req, err := c.preparerForRegistryCodeVersionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryCodeVersionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryCodeVersionsListComplete retrieves all of the results into a single object
func (c CodeVersionClient) RegistryCodeVersionsListComplete(ctx context.Context, id RegistryCodeId, options RegistryCodeVersionsListOperationOptions) (RegistryCodeVersionsListCompleteResult, error) {
	return c.RegistryCodeVersionsListCompleteMatchingPredicate(ctx, id, options, CodeVersionResourceOperationPredicate{})
}

// RegistryCodeVersionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c CodeVersionClient) RegistryCodeVersionsListCompleteMatchingPredicate(ctx context.Context, id RegistryCodeId, options RegistryCodeVersionsListOperationOptions, predicate CodeVersionResourceOperationPredicate) (resp RegistryCodeVersionsListCompleteResult, err error) {
	items := make([]CodeVersionResource, 0)

	page, err := c.RegistryCodeVersionsList(ctx, id, options)
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

	out := RegistryCodeVersionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
