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

type ListWebJobsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWebJobsOperationResponse, error)
}

type ListWebJobsCompleteResult struct {
	Items []WebJob
}

func (r ListWebJobsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWebJobsOperationResponse) LoadMore(ctx context.Context) (resp ListWebJobsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListWebJobs ...
func (c WebAppsClient) ListWebJobs(ctx context.Context, id SiteId) (resp ListWebJobsOperationResponse, err error) {
	req, err := c.preparerForListWebJobs(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWebJobs(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWebJobs prepares the ListWebJobs request.
func (c WebAppsClient) preparerForListWebJobs(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListWebJobsWithNextLink prepares the ListWebJobs request with the given nextLink token.
func (c WebAppsClient) preparerForListWebJobsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWebJobs handles the response to the ListWebJobs request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListWebJobs(resp *http.Response) (result ListWebJobsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWebJobsOperationResponse, err error) {
			req, err := c.preparerForListWebJobsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWebJobs(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWebJobs", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWebJobsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListWebJobsComplete(ctx context.Context, id SiteId) (ListWebJobsCompleteResult, error) {
	return c.ListWebJobsCompleteMatchingPredicate(ctx, id, WebJobOperationPredicate{})
}

// ListWebJobsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListWebJobsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate WebJobOperationPredicate) (resp ListWebJobsCompleteResult, err error) {
	items := make([]WebJob, 0)

	page, err := c.ListWebJobs(ctx, id)
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

	out := ListWebJobsCompleteResult{
		Items: items,
	}
	return out, nil
}
