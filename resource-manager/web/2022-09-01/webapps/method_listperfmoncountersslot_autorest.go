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

type ListPerfMonCountersSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PerfMonResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPerfMonCountersSlotOperationResponse, error)
}

type ListPerfMonCountersSlotCompleteResult struct {
	Items []PerfMonResponse
}

func (r ListPerfMonCountersSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPerfMonCountersSlotOperationResponse) LoadMore(ctx context.Context) (resp ListPerfMonCountersSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListPerfMonCountersSlotOperationOptions struct {
	Filter *string
}

func DefaultListPerfMonCountersSlotOperationOptions() ListPerfMonCountersSlotOperationOptions {
	return ListPerfMonCountersSlotOperationOptions{}
}

func (o ListPerfMonCountersSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListPerfMonCountersSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListPerfMonCountersSlot ...
func (c WebAppsClient) ListPerfMonCountersSlot(ctx context.Context, id SlotId, options ListPerfMonCountersSlotOperationOptions) (resp ListPerfMonCountersSlotOperationResponse, err error) {
	req, err := c.preparerForListPerfMonCountersSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPerfMonCountersSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPerfMonCountersSlot prepares the ListPerfMonCountersSlot request.
func (c WebAppsClient) preparerForListPerfMonCountersSlot(ctx context.Context, id SlotId, options ListPerfMonCountersSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/perfcounters", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListPerfMonCountersSlotWithNextLink prepares the ListPerfMonCountersSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListPerfMonCountersSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPerfMonCountersSlot handles the response to the ListPerfMonCountersSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPerfMonCountersSlot(resp *http.Response) (result ListPerfMonCountersSlotOperationResponse, err error) {
	type page struct {
		Values   []PerfMonResponse `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPerfMonCountersSlotOperationResponse, err error) {
			req, err := c.preparerForListPerfMonCountersSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPerfMonCountersSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPerfMonCountersSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPerfMonCountersSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListPerfMonCountersSlotComplete(ctx context.Context, id SlotId, options ListPerfMonCountersSlotOperationOptions) (ListPerfMonCountersSlotCompleteResult, error) {
	return c.ListPerfMonCountersSlotCompleteMatchingPredicate(ctx, id, options, PerfMonResponseOperationPredicate{})
}

// ListPerfMonCountersSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListPerfMonCountersSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, options ListPerfMonCountersSlotOperationOptions, predicate PerfMonResponseOperationPredicate) (resp ListPerfMonCountersSlotCompleteResult, err error) {
	items := make([]PerfMonResponse, 0)

	page, err := c.ListPerfMonCountersSlot(ctx, id, options)
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

	out := ListPerfMonCountersSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
