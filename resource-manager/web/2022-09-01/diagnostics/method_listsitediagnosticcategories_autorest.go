package diagnostics

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

type ListSiteDiagnosticCategoriesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DiagnosticCategory

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteDiagnosticCategoriesOperationResponse, error)
}

type ListSiteDiagnosticCategoriesCompleteResult struct {
	Items []DiagnosticCategory
}

func (r ListSiteDiagnosticCategoriesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteDiagnosticCategoriesOperationResponse) LoadMore(ctx context.Context) (resp ListSiteDiagnosticCategoriesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteDiagnosticCategories ...
func (c DiagnosticsClient) ListSiteDiagnosticCategories(ctx context.Context, id SiteId) (resp ListSiteDiagnosticCategoriesOperationResponse, err error) {
	req, err := c.preparerForListSiteDiagnosticCategories(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteDiagnosticCategories(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteDiagnosticCategories prepares the ListSiteDiagnosticCategories request.
func (c DiagnosticsClient) preparerForListSiteDiagnosticCategories(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diagnostics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteDiagnosticCategoriesWithNextLink prepares the ListSiteDiagnosticCategories request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteDiagnosticCategoriesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteDiagnosticCategories handles the response to the ListSiteDiagnosticCategories request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteDiagnosticCategories(resp *http.Response) (result ListSiteDiagnosticCategoriesOperationResponse, err error) {
	type page struct {
		Values   []DiagnosticCategory `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteDiagnosticCategoriesOperationResponse, err error) {
			req, err := c.preparerForListSiteDiagnosticCategoriesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteDiagnosticCategories(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDiagnosticCategories", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteDiagnosticCategoriesComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesComplete(ctx context.Context, id SiteId) (ListSiteDiagnosticCategoriesCompleteResult, error) {
	return c.ListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx, id, DiagnosticCategoryOperationPredicate{})
}

// ListSiteDiagnosticCategoriesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate DiagnosticCategoryOperationPredicate) (resp ListSiteDiagnosticCategoriesCompleteResult, err error) {
	items := make([]DiagnosticCategory, 0)

	page, err := c.ListSiteDiagnosticCategories(ctx, id)
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

	out := ListSiteDiagnosticCategoriesCompleteResult{
		Items: items,
	}
	return out, nil
}
