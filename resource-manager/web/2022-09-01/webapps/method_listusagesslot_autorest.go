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

type ListUsagesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmUsageQuota

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListUsagesSlotOperationResponse, error)
}

type ListUsagesSlotCompleteResult struct {
	Items []CsmUsageQuota
}

func (r ListUsagesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListUsagesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListUsagesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListUsagesSlotOperationOptions struct {
	Filter *string
}

func DefaultListUsagesSlotOperationOptions() ListUsagesSlotOperationOptions {
	return ListUsagesSlotOperationOptions{}
}

func (o ListUsagesSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListUsagesSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListUsagesSlot ...
func (c WebAppsClient) ListUsagesSlot(ctx context.Context, id SlotId, options ListUsagesSlotOperationOptions) (resp ListUsagesSlotOperationResponse, err error) {
	req, err := c.preparerForListUsagesSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListUsagesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListUsagesSlot prepares the ListUsagesSlot request.
func (c WebAppsClient) preparerForListUsagesSlot(ctx context.Context, id SlotId, options ListUsagesSlotOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListUsagesSlotWithNextLink prepares the ListUsagesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListUsagesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListUsagesSlot handles the response to the ListUsagesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListUsagesSlot(resp *http.Response) (result ListUsagesSlotOperationResponse, err error) {
	type page struct {
		Values   []CsmUsageQuota `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListUsagesSlotOperationResponse, err error) {
			req, err := c.preparerForListUsagesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListUsagesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListUsagesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListUsagesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListUsagesSlotComplete(ctx context.Context, id SlotId, options ListUsagesSlotOperationOptions) (ListUsagesSlotCompleteResult, error) {
	return c.ListUsagesSlotCompleteMatchingPredicate(ctx, id, options, CsmUsageQuotaOperationPredicate{})
}

// ListUsagesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListUsagesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, options ListUsagesSlotOperationOptions, predicate CsmUsageQuotaOperationPredicate) (resp ListUsagesSlotCompleteResult, err error) {
	items := make([]CsmUsageQuota, 0)

	page, err := c.ListUsagesSlot(ctx, id, options)
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

	out := ListUsagesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
