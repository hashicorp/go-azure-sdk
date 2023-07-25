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

type ListSnapshotsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Snapshot

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSnapshotsSlotOperationResponse, error)
}

type ListSnapshotsSlotCompleteResult struct {
	Items []Snapshot
}

func (r ListSnapshotsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSnapshotsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListSnapshotsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSnapshotsSlot ...
func (c WebAppsClient) ListSnapshotsSlot(ctx context.Context, id SlotId) (resp ListSnapshotsSlotOperationResponse, err error) {
	req, err := c.preparerForListSnapshotsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSnapshotsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSnapshotsSlot prepares the ListSnapshotsSlot request.
func (c WebAppsClient) preparerForListSnapshotsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/snapshots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSnapshotsSlotWithNextLink prepares the ListSnapshotsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListSnapshotsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSnapshotsSlot handles the response to the ListSnapshotsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSnapshotsSlot(resp *http.Response) (result ListSnapshotsSlotOperationResponse, err error) {
	type page struct {
		Values   []Snapshot `json:"value"`
		NextLink *string    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSnapshotsSlotOperationResponse, err error) {
			req, err := c.preparerForListSnapshotsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSnapshotsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSnapshotsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSnapshotsSlotComplete(ctx context.Context, id SlotId) (ListSnapshotsSlotCompleteResult, error) {
	return c.ListSnapshotsSlotCompleteMatchingPredicate(ctx, id, SnapshotOperationPredicate{})
}

// ListSnapshotsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSnapshotsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SnapshotOperationPredicate) (resp ListSnapshotsSlotCompleteResult, err error) {
	items := make([]Snapshot, 0)

	page, err := c.ListSnapshotsSlot(ctx, id)
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

	out := ListSnapshotsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
