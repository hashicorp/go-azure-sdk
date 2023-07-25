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

type ListSnapshotsFromDRSecondarySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Snapshot

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSnapshotsFromDRSecondarySlotOperationResponse, error)
}

type ListSnapshotsFromDRSecondarySlotCompleteResult struct {
	Items []Snapshot
}

func (r ListSnapshotsFromDRSecondarySlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSnapshotsFromDRSecondarySlotOperationResponse) LoadMore(ctx context.Context) (resp ListSnapshotsFromDRSecondarySlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSnapshotsFromDRSecondarySlot ...
func (c WebAppsClient) ListSnapshotsFromDRSecondarySlot(ctx context.Context, id SlotId) (resp ListSnapshotsFromDRSecondarySlotOperationResponse, err error) {
	req, err := c.preparerForListSnapshotsFromDRSecondarySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSnapshotsFromDRSecondarySlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSnapshotsFromDRSecondarySlot prepares the ListSnapshotsFromDRSecondarySlot request.
func (c WebAppsClient) preparerForListSnapshotsFromDRSecondarySlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/snapshotsdr", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSnapshotsFromDRSecondarySlotWithNextLink prepares the ListSnapshotsFromDRSecondarySlot request with the given nextLink token.
func (c WebAppsClient) preparerForListSnapshotsFromDRSecondarySlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSnapshotsFromDRSecondarySlot handles the response to the ListSnapshotsFromDRSecondarySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSnapshotsFromDRSecondarySlot(resp *http.Response) (result ListSnapshotsFromDRSecondarySlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSnapshotsFromDRSecondarySlotOperationResponse, err error) {
			req, err := c.preparerForListSnapshotsFromDRSecondarySlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSnapshotsFromDRSecondarySlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshotsFromDRSecondarySlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSnapshotsFromDRSecondarySlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSnapshotsFromDRSecondarySlotComplete(ctx context.Context, id SlotId) (ListSnapshotsFromDRSecondarySlotCompleteResult, error) {
	return c.ListSnapshotsFromDRSecondarySlotCompleteMatchingPredicate(ctx, id, SnapshotOperationPredicate{})
}

// ListSnapshotsFromDRSecondarySlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSnapshotsFromDRSecondarySlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SnapshotOperationPredicate) (resp ListSnapshotsFromDRSecondarySlotCompleteResult, err error) {
	items := make([]Snapshot, 0)

	page, err := c.ListSnapshotsFromDRSecondarySlot(ctx, id)
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

	out := ListSnapshotsFromDRSecondarySlotCompleteResult{
		Items: items,
	}
	return out, nil
}
