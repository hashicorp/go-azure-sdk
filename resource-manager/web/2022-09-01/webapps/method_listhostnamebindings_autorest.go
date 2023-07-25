package webapps

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

type ListHostNameBindingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HostNameBinding

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListHostNameBindingsOperationResponse, error)
}

type ListHostNameBindingsCompleteResult struct {
	Items []HostNameBinding
}

func (r ListHostNameBindingsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListHostNameBindingsOperationResponse) LoadMore(ctx context.Context) (resp ListHostNameBindingsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListHostNameBindings ...
func (c WebAppsClient) ListHostNameBindings(ctx context.Context, id SiteId) (resp ListHostNameBindingsOperationResponse, err error) {
	req, err := c.preparerForListHostNameBindings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListHostNameBindings(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListHostNameBindings prepares the ListHostNameBindings request.
func (c WebAppsClient) preparerForListHostNameBindings(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hostNameBindings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListHostNameBindingsWithNextLink prepares the ListHostNameBindings request with the given nextLink token.
func (c WebAppsClient) preparerForListHostNameBindingsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListHostNameBindings handles the response to the ListHostNameBindings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHostNameBindings(resp *http.Response) (result ListHostNameBindingsOperationResponse, err error) {
	type page struct {
		Values   []HostNameBinding `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListHostNameBindingsOperationResponse, err error) {
			req, err := c.preparerForListHostNameBindingsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListHostNameBindings(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindings", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListHostNameBindingsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListHostNameBindingsComplete(ctx context.Context, id SiteId) (ListHostNameBindingsCompleteResult, error) {
	return c.ListHostNameBindingsCompleteMatchingPredicate(ctx, id, HostNameBindingOperationPredicate{})
}

// ListHostNameBindingsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListHostNameBindingsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate HostNameBindingOperationPredicate) (resp ListHostNameBindingsCompleteResult, err error) {
	items := make([]HostNameBinding, 0)

	page, err := c.ListHostNameBindings(ctx, id)
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

	out := ListHostNameBindingsCompleteResult{
		Items: items,
	}
	return out, nil
}
