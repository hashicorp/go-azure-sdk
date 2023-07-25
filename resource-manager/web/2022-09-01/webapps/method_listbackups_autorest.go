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

type ListBackupsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BackupItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBackupsOperationResponse, error)
}

type ListBackupsCompleteResult struct {
	Items []BackupItem
}

func (r ListBackupsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBackupsOperationResponse) LoadMore(ctx context.Context) (resp ListBackupsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBackups ...
func (c WebAppsClient) ListBackups(ctx context.Context, id SiteId) (resp ListBackupsOperationResponse, err error) {
	req, err := c.preparerForListBackups(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBackups(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBackups prepares the ListBackups request.
func (c WebAppsClient) preparerForListBackups(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBackupsWithNextLink prepares the ListBackups request with the given nextLink token.
func (c WebAppsClient) preparerForListBackupsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBackups handles the response to the ListBackups request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListBackups(resp *http.Response) (result ListBackupsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBackupsOperationResponse, err error) {
			req, err := c.preparerForListBackupsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBackups(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackups", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBackupsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListBackupsComplete(ctx context.Context, id SiteId) (ListBackupsCompleteResult, error) {
	return c.ListBackupsCompleteMatchingPredicate(ctx, id, BackupItemOperationPredicate{})
}

// ListBackupsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListBackupsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate BackupItemOperationPredicate) (resp ListBackupsCompleteResult, err error) {
	items := make([]BackupItem, 0)

	page, err := c.ListBackups(ctx, id)
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

	out := ListBackupsCompleteResult{
		Items: items,
	}
	return out, nil
}
