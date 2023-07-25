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

type ListTriggeredWebJobHistoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TriggeredJobHistory

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListTriggeredWebJobHistoryOperationResponse, error)
}

type ListTriggeredWebJobHistoryCompleteResult struct {
	Items []TriggeredJobHistory
}

func (r ListTriggeredWebJobHistoryOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListTriggeredWebJobHistoryOperationResponse) LoadMore(ctx context.Context) (resp ListTriggeredWebJobHistoryOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListTriggeredWebJobHistory ...
func (c WebAppsClient) ListTriggeredWebJobHistory(ctx context.Context, id TriggeredWebJobId) (resp ListTriggeredWebJobHistoryOperationResponse, err error) {
	req, err := c.preparerForListTriggeredWebJobHistory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListTriggeredWebJobHistory(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListTriggeredWebJobHistory prepares the ListTriggeredWebJobHistory request.
func (c WebAppsClient) preparerForListTriggeredWebJobHistory(ctx context.Context, id TriggeredWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/history", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListTriggeredWebJobHistoryWithNextLink prepares the ListTriggeredWebJobHistory request with the given nextLink token.
func (c WebAppsClient) preparerForListTriggeredWebJobHistoryWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListTriggeredWebJobHistory handles the response to the ListTriggeredWebJobHistory request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListTriggeredWebJobHistory(resp *http.Response) (result ListTriggeredWebJobHistoryOperationResponse, err error) {
	type page struct {
		Values   []TriggeredJobHistory `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListTriggeredWebJobHistoryOperationResponse, err error) {
			req, err := c.preparerForListTriggeredWebJobHistoryWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListTriggeredWebJobHistory(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListTriggeredWebJobHistory", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListTriggeredWebJobHistoryComplete retrieves all of the results into a single object
func (c WebAppsClient) ListTriggeredWebJobHistoryComplete(ctx context.Context, id TriggeredWebJobId) (ListTriggeredWebJobHistoryCompleteResult, error) {
	return c.ListTriggeredWebJobHistoryCompleteMatchingPredicate(ctx, id, TriggeredJobHistoryOperationPredicate{})
}

// ListTriggeredWebJobHistoryCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListTriggeredWebJobHistoryCompleteMatchingPredicate(ctx context.Context, id TriggeredWebJobId, predicate TriggeredJobHistoryOperationPredicate) (resp ListTriggeredWebJobHistoryCompleteResult, err error) {
	items := make([]TriggeredJobHistory, 0)

	page, err := c.ListTriggeredWebJobHistory(ctx, id)
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

	out := ListTriggeredWebJobHistoryCompleteResult{
		Items: items,
	}
	return out, nil
}
