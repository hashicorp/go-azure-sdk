package staticsites

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

type ListStaticSiteFunctionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteFunctionOverviewARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListStaticSiteFunctionsOperationResponse, error)
}

type ListStaticSiteFunctionsCompleteResult struct {
	Items []StaticSiteFunctionOverviewARMResource
}

func (r ListStaticSiteFunctionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListStaticSiteFunctionsOperationResponse) LoadMore(ctx context.Context) (resp ListStaticSiteFunctionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListStaticSiteFunctions ...
func (c StaticSitesClient) ListStaticSiteFunctions(ctx context.Context, id StaticSiteId) (resp ListStaticSiteFunctionsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteFunctions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListStaticSiteFunctions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListStaticSiteFunctions prepares the ListStaticSiteFunctions request.
func (c StaticSitesClient) preparerForListStaticSiteFunctions(ctx context.Context, id StaticSiteId) (*http.Request, error) {
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

// preparerForListStaticSiteFunctionsWithNextLink prepares the ListStaticSiteFunctions request with the given nextLink token.
func (c StaticSitesClient) preparerForListStaticSiteFunctionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListStaticSiteFunctions handles the response to the ListStaticSiteFunctions request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteFunctions(resp *http.Response) (result ListStaticSiteFunctionsOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteFunctionOverviewARMResource `json:"value"`
		NextLink *string                                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListStaticSiteFunctionsOperationResponse, err error) {
			req, err := c.preparerForListStaticSiteFunctionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListStaticSiteFunctions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteFunctions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListStaticSiteFunctionsComplete retrieves all of the results into a single object
func (c StaticSitesClient) ListStaticSiteFunctionsComplete(ctx context.Context, id StaticSiteId) (ListStaticSiteFunctionsCompleteResult, error) {
	return c.ListStaticSiteFunctionsCompleteMatchingPredicate(ctx, id, StaticSiteFunctionOverviewARMResourceOperationPredicate{})
}

// ListStaticSiteFunctionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) ListStaticSiteFunctionsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteFunctionOverviewARMResourceOperationPredicate) (resp ListStaticSiteFunctionsCompleteResult, err error) {
	items := make([]StaticSiteFunctionOverviewARMResource, 0)

	page, err := c.ListStaticSiteFunctions(ctx, id)
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

	out := ListStaticSiteFunctionsCompleteResult{
		Items: items,
	}
	return out, nil
}
