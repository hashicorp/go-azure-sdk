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

type ListInstanceProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessModuleInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessModulesOperationResponse, error)
}

type ListInstanceProcessModulesCompleteResult struct {
	Items []ProcessModuleInfo
}

func (r ListInstanceProcessModulesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessModulesOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessModulesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcessModules ...
func (c WebAppsClient) ListInstanceProcessModules(ctx context.Context, id InstanceProcessId) (resp ListInstanceProcessModulesOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcessModules(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcessModules(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcessModules prepares the ListInstanceProcessModules request.
func (c WebAppsClient) preparerForListInstanceProcessModules(ctx context.Context, id InstanceProcessId) (*http.Request, error) {
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

// preparerForListInstanceProcessModulesWithNextLink prepares the ListInstanceProcessModules request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessModulesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcessModules handles the response to the ListInstanceProcessModules request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcessModules(resp *http.Response) (result ListInstanceProcessModulesOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessModulesOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessModulesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcessModules(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModules", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessModulesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessModulesComplete(ctx context.Context, id InstanceProcessId) (ListInstanceProcessModulesCompleteResult, error) {
	return c.ListInstanceProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListInstanceProcessModulesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessModulesCompleteMatchingPredicate(ctx context.Context, id InstanceProcessId, predicate ProcessModuleInfoOperationPredicate) (resp ListInstanceProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	page, err := c.ListInstanceProcessModules(ctx, id)
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

	out := ListInstanceProcessModulesCompleteResult{
		Items: items,
	}
	return out, nil
}
