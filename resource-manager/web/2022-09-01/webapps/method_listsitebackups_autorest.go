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

type ListSiteBackupsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BackupItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteBackupsOperationResponse, error)
}

type ListSiteBackupsCompleteResult struct {
	Items []BackupItem
}

func (r ListSiteBackupsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteBackupsOperationResponse) LoadMore(ctx context.Context) (resp ListSiteBackupsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteBackups ...
func (c WebAppsClient) ListSiteBackups(ctx context.Context, id SiteId) (resp ListSiteBackupsOperationResponse, err error) {
	req, err := c.preparerForListSiteBackups(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteBackups(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteBackups prepares the ListSiteBackups request.
func (c WebAppsClient) preparerForListSiteBackups(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listbackups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteBackupsWithNextLink prepares the ListSiteBackups request with the given nextLink token.
func (c WebAppsClient) preparerForListSiteBackupsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSiteBackups handles the response to the ListSiteBackups request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSiteBackups(resp *http.Response) (result ListSiteBackupsOperationResponse, err error) {
	type page struct {
		Values   []BackupItem `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteBackupsOperationResponse, err error) {
			req, err := c.preparerForListSiteBackupsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteBackups(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSiteBackups", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteBackupsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSiteBackupsComplete(ctx context.Context, id SiteId) (ListSiteBackupsCompleteResult, error) {
	return c.ListSiteBackupsCompleteMatchingPredicate(ctx, id, BackupItemOperationPredicate{})
}

// ListSiteBackupsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSiteBackupsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate BackupItemOperationPredicate) (resp ListSiteBackupsCompleteResult, err error) {
	items := make([]BackupItem, 0)

	page, err := c.ListSiteBackups(ctx, id)
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

	out := ListSiteBackupsCompleteResult{
		Items: items,
	}
	return out, nil
}
