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

type ListTriggeredWebJobHistorySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TriggeredJobHistory

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListTriggeredWebJobHistorySlotOperationResponse, error)
}

type ListTriggeredWebJobHistorySlotCompleteResult struct {
	Items []TriggeredJobHistory
}

func (r ListTriggeredWebJobHistorySlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListTriggeredWebJobHistorySlotOperationResponse) LoadMore(ctx context.Context) (resp ListTriggeredWebJobHistorySlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListTriggeredWebJobHistorySlot ...
func (c WebAppsClient) ListTriggeredWebJobHistorySlot(ctx context.Context, id SlotTriggeredWebJobId) (resp ListTriggeredWebJobHistorySlotOperationResponse, err error) {
	req, err := c.preparerForListTriggeredWebJobHistorySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListTriggeredWebJobHistorySlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListTriggeredWebJobHistorySlot prepares the ListTriggeredWebJobHistorySlot request.
func (c WebAppsClient) preparerForListTriggeredWebJobHistorySlot(ctx context.Context, id SlotTriggeredWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/history", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListTriggeredWebJobHistorySlotWithNextLink prepares the ListTriggeredWebJobHistorySlot request with the given nextLink token.
func (c WebAppsClient) preparerForListTriggeredWebJobHistorySlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListTriggeredWebJobHistorySlot handles the response to the ListTriggeredWebJobHistorySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListTriggeredWebJobHistorySlot(resp *http.Response) (result ListTriggeredWebJobHistorySlotOperationResponse, err error) {
	type page struct {
		Values   []TriggeredJobHistory `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListTriggeredWebJobHistorySlotOperationResponse, err error) {
			req, err := c.preparerForListTriggeredWebJobHistorySlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListTriggeredWebJobHistorySlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistorySlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListTriggeredWebJobHistorySlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListTriggeredWebJobHistorySlotComplete(ctx context.Context, id SlotTriggeredWebJobId) (ListTriggeredWebJobHistorySlotCompleteResult, error) {
	return c.ListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx, id, TriggeredJobHistoryOperationPredicate{})
}

// ListTriggeredWebJobHistorySlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx context.Context, id SlotTriggeredWebJobId, predicate TriggeredJobHistoryOperationPredicate) (resp ListTriggeredWebJobHistorySlotCompleteResult, err error) {
	items := make([]TriggeredJobHistory, 0)

	page, err := c.ListTriggeredWebJobHistorySlot(ctx, id)
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

	out := ListTriggeredWebJobHistorySlotCompleteResult{
		Items: items,
	}
	return out, nil
}
