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

type ListTriggeredWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TriggeredWebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListTriggeredWebJobsSlotOperationResponse, error)
}

type ListTriggeredWebJobsSlotCompleteResult struct {
	Items []TriggeredWebJob
}

func (r ListTriggeredWebJobsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListTriggeredWebJobsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListTriggeredWebJobsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListTriggeredWebJobsSlot ...
func (c WebAppsClient) ListTriggeredWebJobsSlot(ctx context.Context, id SlotId) (resp ListTriggeredWebJobsSlotOperationResponse, err error) {
	req, err := c.preparerForListTriggeredWebJobsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListTriggeredWebJobsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListTriggeredWebJobsSlot prepares the ListTriggeredWebJobsSlot request.
func (c WebAppsClient) preparerForListTriggeredWebJobsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/triggeredWebJobs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListTriggeredWebJobsSlotWithNextLink prepares the ListTriggeredWebJobsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListTriggeredWebJobsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListTriggeredWebJobsSlot handles the response to the ListTriggeredWebJobsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListTriggeredWebJobsSlot(resp *http.Response) (result ListTriggeredWebJobsSlotOperationResponse, err error) {
	type page struct {
		Values   []TriggeredWebJob `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListTriggeredWebJobsSlotOperationResponse, err error) {
			req, err := c.preparerForListTriggeredWebJobsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListTriggeredWebJobsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListTriggeredWebJobsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListTriggeredWebJobsSlotComplete(ctx context.Context, id SlotId) (ListTriggeredWebJobsSlotCompleteResult, error) {
	return c.ListTriggeredWebJobsSlotCompleteMatchingPredicate(ctx, id, TriggeredWebJobOperationPredicate{})
}

// ListTriggeredWebJobsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListTriggeredWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate TriggeredWebJobOperationPredicate) (resp ListTriggeredWebJobsSlotCompleteResult, err error) {
	items := make([]TriggeredWebJob, 0)

	page, err := c.ListTriggeredWebJobsSlot(ctx, id)
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

	out := ListTriggeredWebJobsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
