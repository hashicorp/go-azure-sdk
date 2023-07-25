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

type ListSlotDifferencesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SlotDifference

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSlotDifferencesSlotOperationResponse, error)
}

type ListSlotDifferencesSlotCompleteResult struct {
	Items []SlotDifference
}

func (r ListSlotDifferencesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSlotDifferencesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListSlotDifferencesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSlotDifferencesSlot ...
func (c WebAppsClient) ListSlotDifferencesSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (resp ListSlotDifferencesSlotOperationResponse, err error) {
	req, err := c.preparerForListSlotDifferencesSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSlotDifferencesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSlotDifferencesSlot prepares the ListSlotDifferencesSlot request.
func (c WebAppsClient) preparerForListSlotDifferencesSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/slotsdiffs", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSlotDifferencesSlotWithNextLink prepares the ListSlotDifferencesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListSlotDifferencesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSlotDifferencesSlot handles the response to the ListSlotDifferencesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSlotDifferencesSlot(resp *http.Response) (result ListSlotDifferencesSlotOperationResponse, err error) {
	type page struct {
		Values   []SlotDifference `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSlotDifferencesSlotOperationResponse, err error) {
			req, err := c.preparerForListSlotDifferencesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSlotDifferencesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSlotDifferencesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSlotDifferencesSlotComplete(ctx context.Context, id SlotId, input CsmSlotEntity) (ListSlotDifferencesSlotCompleteResult, error) {
	return c.ListSlotDifferencesSlotCompleteMatchingPredicate(ctx, id, input, SlotDifferenceOperationPredicate{})
}

// ListSlotDifferencesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSlotDifferencesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, input CsmSlotEntity, predicate SlotDifferenceOperationPredicate) (resp ListSlotDifferencesSlotCompleteResult, err error) {
	items := make([]SlotDifference, 0)

	page, err := c.ListSlotDifferencesSlot(ctx, id, input)
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

	out := ListSlotDifferencesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
