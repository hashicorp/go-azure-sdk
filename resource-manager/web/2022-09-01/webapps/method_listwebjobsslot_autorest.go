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

type ListWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWebJobsSlotOperationResponse, error)
}

type ListWebJobsSlotCompleteResult struct {
	Items []WebJob
}

func (r ListWebJobsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWebJobsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListWebJobsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListWebJobsSlot ...
func (c WebAppsClient) ListWebJobsSlot(ctx context.Context, id SlotId) (resp ListWebJobsSlotOperationResponse, err error) {
	req, err := c.preparerForListWebJobsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWebJobsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWebJobsSlot prepares the ListWebJobsSlot request.
func (c WebAppsClient) preparerForListWebJobsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/webJobs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWebJobsSlotWithNextLink prepares the ListWebJobsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListWebJobsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWebJobsSlot handles the response to the ListWebJobsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListWebJobsSlot(resp *http.Response) (result ListWebJobsSlotOperationResponse, err error) {
	type page struct {
		Values   []WebJob `json:"value"`
		NextLink *string  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWebJobsSlotOperationResponse, err error) {
			req, err := c.preparerForListWebJobsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWebJobsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWebJobsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListWebJobsSlotComplete(ctx context.Context, id SlotId) (ListWebJobsSlotCompleteResult, error) {
	return c.ListWebJobsSlotCompleteMatchingPredicate(ctx, id, WebJobOperationPredicate{})
}

// ListWebJobsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WebJobOperationPredicate) (resp ListWebJobsSlotCompleteResult, err error) {
	items := make([]WebJob, 0)

	page, err := c.ListWebJobsSlot(ctx, id)
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

	out := ListWebJobsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
