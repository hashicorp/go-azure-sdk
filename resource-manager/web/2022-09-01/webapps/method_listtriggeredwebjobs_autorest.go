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

type ListTriggeredWebJobsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TriggeredWebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListTriggeredWebJobsOperationResponse, error)
}

type ListTriggeredWebJobsCompleteResult struct {
	Items []TriggeredWebJob
}

func (r ListTriggeredWebJobsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListTriggeredWebJobsOperationResponse) LoadMore(ctx context.Context) (resp ListTriggeredWebJobsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListTriggeredWebJobs ...
func (c WebAppsClient) ListTriggeredWebJobs(ctx context.Context, id SiteId) (resp ListTriggeredWebJobsOperationResponse, err error) {
	req, err := c.preparerForListTriggeredWebJobs(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListTriggeredWebJobs(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListTriggeredWebJobs prepares the ListTriggeredWebJobs request.
func (c WebAppsClient) preparerForListTriggeredWebJobs(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListTriggeredWebJobsWithNextLink prepares the ListTriggeredWebJobs request with the given nextLink token.
func (c WebAppsClient) preparerForListTriggeredWebJobsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListTriggeredWebJobs handles the response to the ListTriggeredWebJobs request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListTriggeredWebJobs(resp *http.Response) (result ListTriggeredWebJobsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListTriggeredWebJobsOperationResponse, err error) {
			req, err := c.preparerForListTriggeredWebJobsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListTriggeredWebJobs(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobs", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListTriggeredWebJobsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListTriggeredWebJobsComplete(ctx context.Context, id SiteId) (ListTriggeredWebJobsCompleteResult, error) {
	return c.ListTriggeredWebJobsCompleteMatchingPredicate(ctx, id, TriggeredWebJobOperationPredicate{})
}

// ListTriggeredWebJobsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListTriggeredWebJobsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate TriggeredWebJobOperationPredicate) (resp ListTriggeredWebJobsCompleteResult, err error) {
	items := make([]TriggeredWebJob, 0)

	page, err := c.ListTriggeredWebJobs(ctx, id)
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

	out := ListTriggeredWebJobsCompleteResult{
		Items: items,
	}
	return out, nil
}
