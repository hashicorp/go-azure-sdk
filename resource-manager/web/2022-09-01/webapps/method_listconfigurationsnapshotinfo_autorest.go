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

type ListConfigurationSnapshotInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SiteConfigurationSnapshotInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListConfigurationSnapshotInfoOperationResponse, error)
}

type ListConfigurationSnapshotInfoCompleteResult struct {
	Items []SiteConfigurationSnapshotInfo
}

func (r ListConfigurationSnapshotInfoOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListConfigurationSnapshotInfoOperationResponse) LoadMore(ctx context.Context) (resp ListConfigurationSnapshotInfoOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListConfigurationSnapshotInfo ...
func (c WebAppsClient) ListConfigurationSnapshotInfo(ctx context.Context, id SiteId) (resp ListConfigurationSnapshotInfoOperationResponse, err error) {
	req, err := c.preparerForListConfigurationSnapshotInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListConfigurationSnapshotInfo(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListConfigurationSnapshotInfo prepares the ListConfigurationSnapshotInfo request.
func (c WebAppsClient) preparerForListConfigurationSnapshotInfo(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web/snapshots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListConfigurationSnapshotInfoWithNextLink prepares the ListConfigurationSnapshotInfo request with the given nextLink token.
func (c WebAppsClient) preparerForListConfigurationSnapshotInfoWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListConfigurationSnapshotInfo handles the response to the ListConfigurationSnapshotInfo request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListConfigurationSnapshotInfo(resp *http.Response) (result ListConfigurationSnapshotInfoOperationResponse, err error) {
	type page struct {
		Values   []SiteConfigurationSnapshotInfo `json:"value"`
		NextLink *string                         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListConfigurationSnapshotInfoOperationResponse, err error) {
			req, err := c.preparerForListConfigurationSnapshotInfoWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListConfigurationSnapshotInfo(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationSnapshotInfo", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListConfigurationSnapshotInfoComplete retrieves all of the results into a single object
func (c WebAppsClient) ListConfigurationSnapshotInfoComplete(ctx context.Context, id SiteId) (ListConfigurationSnapshotInfoCompleteResult, error) {
	return c.ListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx, id, SiteConfigurationSnapshotInfoOperationPredicate{})
}

// ListConfigurationSnapshotInfoCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate SiteConfigurationSnapshotInfoOperationPredicate) (resp ListConfigurationSnapshotInfoCompleteResult, err error) {
	items := make([]SiteConfigurationSnapshotInfo, 0)

	page, err := c.ListConfigurationSnapshotInfo(ctx, id)
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

	out := ListConfigurationSnapshotInfoCompleteResult{
		Items: items,
	}
	return out, nil
}
