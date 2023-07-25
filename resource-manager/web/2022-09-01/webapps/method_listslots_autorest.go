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

type ListSlotsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Site

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSlotsOperationResponse, error)
}

type ListSlotsCompleteResult struct {
	Items []Site
}

func (r ListSlotsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSlotsOperationResponse) LoadMore(ctx context.Context) (resp ListSlotsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSlots ...
func (c WebAppsClient) ListSlots(ctx context.Context, id SiteId) (resp ListSlotsOperationResponse, err error) {
	req, err := c.preparerForListSlots(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSlots(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSlots prepares the ListSlots request.
func (c WebAppsClient) preparerForListSlots(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/slots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSlotsWithNextLink prepares the ListSlots request with the given nextLink token.
func (c WebAppsClient) preparerForListSlotsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSlots handles the response to the ListSlots request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSlots(resp *http.Response) (result ListSlotsOperationResponse, err error) {
	type page struct {
		Values   []Site  `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSlotsOperationResponse, err error) {
			req, err := c.preparerForListSlotsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSlots(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlots", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSlotsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSlotsComplete(ctx context.Context, id SiteId) (ListSlotsCompleteResult, error) {
	return c.ListSlotsCompleteMatchingPredicate(ctx, id, SiteOperationPredicate{})
}

// ListSlotsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSlotsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate SiteOperationPredicate) (resp ListSlotsCompleteResult, err error) {
	items := make([]Site, 0)

	page, err := c.ListSlots(ctx, id)
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

	out := ListSlotsCompleteResult{
		Items: items,
	}
	return out, nil
}
