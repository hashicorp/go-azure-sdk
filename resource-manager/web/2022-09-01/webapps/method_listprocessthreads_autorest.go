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

type ListProcessThreadsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessThreadInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProcessThreadsOperationResponse, error)
}

type ListProcessThreadsCompleteResult struct {
	Items []ProcessThreadInfo
}

func (r ListProcessThreadsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProcessThreadsOperationResponse) LoadMore(ctx context.Context) (resp ListProcessThreadsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProcessThreads ...
func (c WebAppsClient) ListProcessThreads(ctx context.Context, id ProcessId) (resp ListProcessThreadsOperationResponse, err error) {
	req, err := c.preparerForListProcessThreads(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProcessThreads(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProcessThreads prepares the ListProcessThreads request.
func (c WebAppsClient) preparerForListProcessThreads(ctx context.Context, id ProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/threads", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListProcessThreadsWithNextLink prepares the ListProcessThreads request with the given nextLink token.
func (c WebAppsClient) preparerForListProcessThreadsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProcessThreads handles the response to the ListProcessThreads request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProcessThreads(resp *http.Response) (result ListProcessThreadsOperationResponse, err error) {
	type page struct {
		Values   []ProcessThreadInfo `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProcessThreadsOperationResponse, err error) {
			req, err := c.preparerForListProcessThreadsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProcessThreads(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessThreads", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProcessThreadsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProcessThreadsComplete(ctx context.Context, id ProcessId) (ListProcessThreadsCompleteResult, error) {
	return c.ListProcessThreadsCompleteMatchingPredicate(ctx, id, ProcessThreadInfoOperationPredicate{})
}

// ListProcessThreadsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProcessThreadsCompleteMatchingPredicate(ctx context.Context, id ProcessId, predicate ProcessThreadInfoOperationPredicate) (resp ListProcessThreadsCompleteResult, err error) {
	items := make([]ProcessThreadInfo, 0)

	page, err := c.ListProcessThreads(ctx, id)
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

	out := ListProcessThreadsCompleteResult{
		Items: items,
	}
	return out, nil
}
