package appserviceplans

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

type ListHybridConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HybridConnection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListHybridConnectionsOperationResponse, error)
}

type ListHybridConnectionsCompleteResult struct {
	Items []HybridConnection
}

func (r ListHybridConnectionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListHybridConnectionsOperationResponse) LoadMore(ctx context.Context) (resp ListHybridConnectionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListHybridConnections ...
func (c AppServicePlansClient) ListHybridConnections(ctx context.Context, id ServerFarmId) (resp ListHybridConnectionsOperationResponse, err error) {
	req, err := c.preparerForListHybridConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListHybridConnections(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListHybridConnections prepares the ListHybridConnections request.
func (c AppServicePlansClient) preparerForListHybridConnections(ctx context.Context, id ServerFarmId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnectionRelays", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListHybridConnectionsWithNextLink prepares the ListHybridConnections request with the given nextLink token.
func (c AppServicePlansClient) preparerForListHybridConnectionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListHybridConnections handles the response to the ListHybridConnections request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListHybridConnections(resp *http.Response) (result ListHybridConnectionsOperationResponse, err error) {
	type page struct {
		Values   []HybridConnection `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListHybridConnectionsOperationResponse, err error) {
			req, err := c.preparerForListHybridConnectionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListHybridConnections(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnections", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListHybridConnectionsComplete retrieves all of the results into a single object
func (c AppServicePlansClient) ListHybridConnectionsComplete(ctx context.Context, id ServerFarmId) (ListHybridConnectionsCompleteResult, error) {
	return c.ListHybridConnectionsCompleteMatchingPredicate(ctx, id, HybridConnectionOperationPredicate{})
}

// ListHybridConnectionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServicePlansClient) ListHybridConnectionsCompleteMatchingPredicate(ctx context.Context, id ServerFarmId, predicate HybridConnectionOperationPredicate) (resp ListHybridConnectionsCompleteResult, err error) {
	items := make([]HybridConnection, 0)

	page, err := c.ListHybridConnections(ctx, id)
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

	out := ListHybridConnectionsCompleteResult{
		Items: items,
	}
	return out, nil
}
