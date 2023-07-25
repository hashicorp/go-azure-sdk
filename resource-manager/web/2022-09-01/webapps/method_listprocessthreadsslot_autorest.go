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

type ListProcessThreadsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessThreadInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProcessThreadsSlotOperationResponse, error)
}

type ListProcessThreadsSlotCompleteResult struct {
	Items []ProcessThreadInfo
}

func (r ListProcessThreadsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProcessThreadsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListProcessThreadsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProcessThreadsSlot ...
func (c WebAppsClient) ListProcessThreadsSlot(ctx context.Context, id SlotProcessId) (resp ListProcessThreadsSlotOperationResponse, err error) {
	req, err := c.preparerForListProcessThreadsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProcessThreadsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProcessThreadsSlot prepares the ListProcessThreadsSlot request.
func (c WebAppsClient) preparerForListProcessThreadsSlot(ctx context.Context, id SlotProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/threads", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListProcessThreadsSlotWithNextLink prepares the ListProcessThreadsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListProcessThreadsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProcessThreadsSlot handles the response to the ListProcessThreadsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProcessThreadsSlot(resp *http.Response) (result ListProcessThreadsSlotOperationResponse, err error) {
	type page struct {
		Values   []ProcessThreadInfo `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProcessThreadsSlotOperationResponse, err error) {
			req, err := c.preparerForListProcessThreadsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProcessThreadsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreadsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProcessThreadsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProcessThreadsSlotComplete(ctx context.Context, id SlotProcessId) (ListProcessThreadsSlotCompleteResult, error) {
	return c.ListProcessThreadsSlotCompleteMatchingPredicate(ctx, id, ProcessThreadInfoOperationPredicate{})
}

// ListProcessThreadsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProcessThreadsSlotCompleteMatchingPredicate(ctx context.Context, id SlotProcessId, predicate ProcessThreadInfoOperationPredicate) (resp ListProcessThreadsSlotCompleteResult, err error) {
	items := make([]ProcessThreadInfo, 0)

	page, err := c.ListProcessThreadsSlot(ctx, id)
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

	out := ListProcessThreadsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
