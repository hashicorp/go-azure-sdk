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

type ListInstanceProcessThreadsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessThreadInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessThreadsOperationResponse, error)
}

type ListInstanceProcessThreadsCompleteResult struct {
	Items []ProcessThreadInfo
}

func (r ListInstanceProcessThreadsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessThreadsOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessThreadsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcessThreads ...
func (c WebAppsClient) ListInstanceProcessThreads(ctx context.Context, id InstanceProcessId) (resp ListInstanceProcessThreadsOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcessThreads(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcessThreads(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcessThreads prepares the ListInstanceProcessThreads request.
func (c WebAppsClient) preparerForListInstanceProcessThreads(ctx context.Context, id InstanceProcessId) (*http.Request, error) {
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

// preparerForListInstanceProcessThreadsWithNextLink prepares the ListInstanceProcessThreads request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessThreadsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcessThreads handles the response to the ListInstanceProcessThreads request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcessThreads(resp *http.Response) (result ListInstanceProcessThreadsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessThreadsOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessThreadsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcessThreads(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessThreads", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessThreadsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessThreadsComplete(ctx context.Context, id InstanceProcessId) (ListInstanceProcessThreadsCompleteResult, error) {
	return c.ListInstanceProcessThreadsCompleteMatchingPredicate(ctx, id, ProcessThreadInfoOperationPredicate{})
}

// ListInstanceProcessThreadsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessThreadsCompleteMatchingPredicate(ctx context.Context, id InstanceProcessId, predicate ProcessThreadInfoOperationPredicate) (resp ListInstanceProcessThreadsCompleteResult, err error) {
	items := make([]ProcessThreadInfo, 0)

	page, err := c.ListInstanceProcessThreads(ctx, id)
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

	out := ListInstanceProcessThreadsCompleteResult{
		Items: items,
	}
	return out, nil
}
