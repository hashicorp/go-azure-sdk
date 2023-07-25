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

type ListContinuousWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ContinuousWebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListContinuousWebJobsSlotOperationResponse, error)
}

type ListContinuousWebJobsSlotCompleteResult struct {
	Items []ContinuousWebJob
}

func (r ListContinuousWebJobsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListContinuousWebJobsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListContinuousWebJobsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListContinuousWebJobsSlot ...
func (c WebAppsClient) ListContinuousWebJobsSlot(ctx context.Context, id SlotId) (resp ListContinuousWebJobsSlotOperationResponse, err error) {
	req, err := c.preparerForListContinuousWebJobsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListContinuousWebJobsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListContinuousWebJobsSlot prepares the ListContinuousWebJobsSlot request.
func (c WebAppsClient) preparerForListContinuousWebJobsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/continuousWebJobs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListContinuousWebJobsSlotWithNextLink prepares the ListContinuousWebJobsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListContinuousWebJobsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListContinuousWebJobsSlot handles the response to the ListContinuousWebJobsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListContinuousWebJobsSlot(resp *http.Response) (result ListContinuousWebJobsSlotOperationResponse, err error) {
	type page struct {
		Values   []ContinuousWebJob `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListContinuousWebJobsSlotOperationResponse, err error) {
			req, err := c.preparerForListContinuousWebJobsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListContinuousWebJobsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListContinuousWebJobsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListContinuousWebJobsSlotComplete(ctx context.Context, id SlotId) (ListContinuousWebJobsSlotCompleteResult, error) {
	return c.ListContinuousWebJobsSlotCompleteMatchingPredicate(ctx, id, ContinuousWebJobOperationPredicate{})
}

// ListContinuousWebJobsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListContinuousWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ContinuousWebJobOperationPredicate) (resp ListContinuousWebJobsSlotCompleteResult, err error) {
	items := make([]ContinuousWebJob, 0)

	page, err := c.ListContinuousWebJobsSlot(ctx, id)
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

	out := ListContinuousWebJobsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
