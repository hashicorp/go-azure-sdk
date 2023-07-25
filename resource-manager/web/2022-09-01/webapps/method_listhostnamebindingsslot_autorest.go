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

type ListHostNameBindingsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HostNameBinding

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListHostNameBindingsSlotOperationResponse, error)
}

type ListHostNameBindingsSlotCompleteResult struct {
	Items []HostNameBinding
}

func (r ListHostNameBindingsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListHostNameBindingsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListHostNameBindingsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListHostNameBindingsSlot ...
func (c WebAppsClient) ListHostNameBindingsSlot(ctx context.Context, id SlotId) (resp ListHostNameBindingsSlotOperationResponse, err error) {
	req, err := c.preparerForListHostNameBindingsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListHostNameBindingsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListHostNameBindingsSlot prepares the ListHostNameBindingsSlot request.
func (c WebAppsClient) preparerForListHostNameBindingsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// preparerForListHostNameBindingsSlotWithNextLink prepares the ListHostNameBindingsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListHostNameBindingsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListHostNameBindingsSlot handles the response to the ListHostNameBindingsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHostNameBindingsSlot(resp *http.Response) (result ListHostNameBindingsSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListHostNameBindingsSlotOperationResponse, err error) {
			req, err := c.preparerForListHostNameBindingsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListHostNameBindingsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostNameBindingsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListHostNameBindingsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListHostNameBindingsSlotComplete(ctx context.Context, id SlotId) (ListHostNameBindingsSlotCompleteResult, error) {
	return c.ListHostNameBindingsSlotCompleteMatchingPredicate(ctx, id, HostNameBindingOperationPredicate{})
}

// ListHostNameBindingsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListHostNameBindingsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate HostNameBindingOperationPredicate) (resp ListHostNameBindingsSlotCompleteResult, err error) {
	items := make([]HostNameBinding, 0)

	page, err := c.ListHostNameBindingsSlot(ctx, id)
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

	out := ListHostNameBindingsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
