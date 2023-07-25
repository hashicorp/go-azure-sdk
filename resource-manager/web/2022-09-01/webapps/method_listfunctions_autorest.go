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

type ListFunctionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FunctionEnvelope

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListFunctionsOperationResponse, error)
}

type ListFunctionsCompleteResult struct {
	Items []FunctionEnvelope
}

func (r ListFunctionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListFunctionsOperationResponse) LoadMore(ctx context.Context) (resp ListFunctionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListFunctions ...
func (c WebAppsClient) ListFunctions(ctx context.Context, id SiteId) (resp ListFunctionsOperationResponse, err error) {
	req, err := c.preparerForListFunctions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListFunctions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListFunctions prepares the ListFunctions request.
func (c WebAppsClient) preparerForListFunctions(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/functions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListFunctionsWithNextLink prepares the ListFunctions request with the given nextLink token.
func (c WebAppsClient) preparerForListFunctionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListFunctions handles the response to the ListFunctions request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListFunctions(resp *http.Response) (result ListFunctionsOperationResponse, err error) {
	type page struct {
		Values   []FunctionEnvelope `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListFunctionsOperationResponse, err error) {
			req, err := c.preparerForListFunctionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListFunctions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListFunctionsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListFunctionsComplete(ctx context.Context, id SiteId) (ListFunctionsCompleteResult, error) {
	return c.ListFunctionsCompleteMatchingPredicate(ctx, id, FunctionEnvelopeOperationPredicate{})
}

// ListFunctionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListFunctionsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate FunctionEnvelopeOperationPredicate) (resp ListFunctionsCompleteResult, err error) {
	items := make([]FunctionEnvelope, 0)

	page, err := c.ListFunctions(ctx, id)
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

	out := ListFunctionsCompleteResult{
		Items: items,
	}
	return out, nil
}
