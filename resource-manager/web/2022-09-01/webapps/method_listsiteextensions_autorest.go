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

type ListSiteExtensionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SiteExtensionInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteExtensionsOperationResponse, error)
}

type ListSiteExtensionsCompleteResult struct {
	Items []SiteExtensionInfo
}

func (r ListSiteExtensionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteExtensionsOperationResponse) LoadMore(ctx context.Context) (resp ListSiteExtensionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteExtensions ...
func (c WebAppsClient) ListSiteExtensions(ctx context.Context, id SiteId) (resp ListSiteExtensionsOperationResponse, err error) {
	req, err := c.preparerForListSiteExtensions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteExtensions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteExtensions prepares the ListSiteExtensions request.
func (c WebAppsClient) preparerForListSiteExtensions(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/siteExtensions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteExtensionsWithNextLink prepares the ListSiteExtensions request with the given nextLink token.
func (c WebAppsClient) preparerForListSiteExtensionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteExtensions handles the response to the ListSiteExtensions request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSiteExtensions(resp *http.Response) (result ListSiteExtensionsOperationResponse, err error) {
	type page struct {
		Values   []SiteExtensionInfo `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteExtensionsOperationResponse, err error) {
			req, err := c.preparerForListSiteExtensionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteExtensions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteExtensions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteExtensionsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSiteExtensionsComplete(ctx context.Context, id SiteId) (ListSiteExtensionsCompleteResult, error) {
	return c.ListSiteExtensionsCompleteMatchingPredicate(ctx, id, SiteExtensionInfoOperationPredicate{})
}

// ListSiteExtensionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSiteExtensionsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate SiteExtensionInfoOperationPredicate) (resp ListSiteExtensionsCompleteResult, err error) {
	items := make([]SiteExtensionInfo, 0)

	page, err := c.ListSiteExtensions(ctx, id)
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

	out := ListSiteExtensionsCompleteResult{
		Items: items,
	}
	return out, nil
}
