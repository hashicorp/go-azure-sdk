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

type ListInstanceProcessThreadsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessThreadInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessThreadsSlotOperationResponse, error)
}

type ListInstanceProcessThreadsSlotCompleteResult struct {
	Items []ProcessThreadInfo
}

func (r ListInstanceProcessThreadsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessThreadsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessThreadsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcessThreadsSlot ...
func (c WebAppsClient) ListInstanceProcessThreadsSlot(ctx context.Context, id SlotInstanceProcessId) (resp ListInstanceProcessThreadsSlotOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcessThreadsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcessThreadsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcessThreadsSlot prepares the ListInstanceProcessThreadsSlot request.
func (c WebAppsClient) preparerForListInstanceProcessThreadsSlot(ctx context.Context, id SlotInstanceProcessId) (*http.Request, error) {
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

// preparerForListInstanceProcessThreadsSlotWithNextLink prepares the ListInstanceProcessThreadsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessThreadsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcessThreadsSlot handles the response to the ListInstanceProcessThreadsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcessThreadsSlot(resp *http.Response) (result ListInstanceProcessThreadsSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessThreadsSlotOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessThreadsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcessThreadsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreadsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessThreadsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessThreadsSlotComplete(ctx context.Context, id SlotInstanceProcessId) (ListInstanceProcessThreadsSlotCompleteResult, error) {
	return c.ListInstanceProcessThreadsSlotCompleteMatchingPredicate(ctx, id, ProcessThreadInfoOperationPredicate{})
}

// ListInstanceProcessThreadsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessThreadsSlotCompleteMatchingPredicate(ctx context.Context, id SlotInstanceProcessId, predicate ProcessThreadInfoOperationPredicate) (resp ListInstanceProcessThreadsSlotCompleteResult, err error) {
	items := make([]ProcessThreadInfo, 0)

	page, err := c.ListInstanceProcessThreadsSlot(ctx, id)
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

	out := ListInstanceProcessThreadsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
