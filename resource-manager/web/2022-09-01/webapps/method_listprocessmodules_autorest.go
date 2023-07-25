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

type ListProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessModuleInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProcessModulesOperationResponse, error)
}

type ListProcessModulesCompleteResult struct {
	Items []ProcessModuleInfo
}

func (r ListProcessModulesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProcessModulesOperationResponse) LoadMore(ctx context.Context) (resp ListProcessModulesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProcessModules ...
func (c WebAppsClient) ListProcessModules(ctx context.Context, id ProcessId) (resp ListProcessModulesOperationResponse, err error) {
	req, err := c.preparerForListProcessModules(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProcessModules(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProcessModules prepares the ListProcessModules request.
func (c WebAppsClient) preparerForListProcessModules(ctx context.Context, id ProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/modules", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListProcessModulesWithNextLink prepares the ListProcessModules request with the given nextLink token.
func (c WebAppsClient) preparerForListProcessModulesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProcessModules handles the response to the ListProcessModules request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProcessModules(resp *http.Response) (result ListProcessModulesOperationResponse, err error) {
	type page struct {
		Values   []ProcessModuleInfo `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProcessModulesOperationResponse, err error) {
			req, err := c.preparerForListProcessModulesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProcessModules(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModules", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProcessModulesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProcessModulesComplete(ctx context.Context, id ProcessId) (ListProcessModulesCompleteResult, error) {
	return c.ListProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListProcessModulesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProcessModulesCompleteMatchingPredicate(ctx context.Context, id ProcessId, predicate ProcessModuleInfoOperationPredicate) (resp ListProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	page, err := c.ListProcessModules(ctx, id)
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

	out := ListProcessModulesCompleteResult{
		Items: items,
	}
	return out, nil
}
