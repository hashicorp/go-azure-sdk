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

type ListInstanceIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebSiteInstanceStatus

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceIdentifiersOperationResponse, error)
}

type ListInstanceIdentifiersCompleteResult struct {
	Items []WebSiteInstanceStatus
}

func (r ListInstanceIdentifiersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceIdentifiersOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceIdentifiersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceIdentifiers ...
func (c WebAppsClient) ListInstanceIdentifiers(ctx context.Context, id SiteId) (resp ListInstanceIdentifiersOperationResponse, err error) {
	req, err := c.preparerForListInstanceIdentifiers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceIdentifiers(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceIdentifiers prepares the ListInstanceIdentifiers request.
func (c WebAppsClient) preparerForListInstanceIdentifiers(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/instances", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListInstanceIdentifiersWithNextLink prepares the ListInstanceIdentifiers request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceIdentifiersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceIdentifiers handles the response to the ListInstanceIdentifiers request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceIdentifiers(resp *http.Response) (result ListInstanceIdentifiersOperationResponse, err error) {
	type page struct {
		Values   []WebSiteInstanceStatus `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceIdentifiersOperationResponse, err error) {
			req, err := c.preparerForListInstanceIdentifiersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceIdentifiers(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiers", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceIdentifiersComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceIdentifiersComplete(ctx context.Context, id SiteId) (ListInstanceIdentifiersCompleteResult, error) {
	return c.ListInstanceIdentifiersCompleteMatchingPredicate(ctx, id, WebSiteInstanceStatusOperationPredicate{})
}

// ListInstanceIdentifiersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceIdentifiersCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate WebSiteInstanceStatusOperationPredicate) (resp ListInstanceIdentifiersCompleteResult, err error) {
	items := make([]WebSiteInstanceStatus, 0)

	page, err := c.ListInstanceIdentifiers(ctx, id)
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

	out := ListInstanceIdentifiersCompleteResult{
		Items: items,
	}
	return out, nil
}
