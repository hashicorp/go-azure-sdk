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

type ListContinuousWebJobsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ContinuousWebJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListContinuousWebJobsOperationResponse, error)
}

type ListContinuousWebJobsCompleteResult struct {
	Items []ContinuousWebJob
}

func (r ListContinuousWebJobsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListContinuousWebJobsOperationResponse) LoadMore(ctx context.Context) (resp ListContinuousWebJobsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListContinuousWebJobs ...
func (c WebAppsClient) ListContinuousWebJobs(ctx context.Context, id SiteId) (resp ListContinuousWebJobsOperationResponse, err error) {
	req, err := c.preparerForListContinuousWebJobs(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListContinuousWebJobs(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListContinuousWebJobs prepares the ListContinuousWebJobs request.
func (c WebAppsClient) preparerForListContinuousWebJobs(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListContinuousWebJobsWithNextLink prepares the ListContinuousWebJobs request with the given nextLink token.
func (c WebAppsClient) preparerForListContinuousWebJobsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListContinuousWebJobs handles the response to the ListContinuousWebJobs request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListContinuousWebJobs(resp *http.Response) (result ListContinuousWebJobsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListContinuousWebJobsOperationResponse, err error) {
			req, err := c.preparerForListContinuousWebJobsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListContinuousWebJobs(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListContinuousWebJobs", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListContinuousWebJobsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListContinuousWebJobsComplete(ctx context.Context, id SiteId) (ListContinuousWebJobsCompleteResult, error) {
	return c.ListContinuousWebJobsCompleteMatchingPredicate(ctx, id, ContinuousWebJobOperationPredicate{})
}

// ListContinuousWebJobsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListContinuousWebJobsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate ContinuousWebJobOperationPredicate) (resp ListContinuousWebJobsCompleteResult, err error) {
	items := make([]ContinuousWebJob, 0)

	page, err := c.ListContinuousWebJobs(ctx, id)
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

	out := ListContinuousWebJobsCompleteResult{
		Items: items,
	}
	return out, nil
}
