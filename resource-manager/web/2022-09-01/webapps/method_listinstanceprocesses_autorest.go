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

type ListInstanceProcessesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessesOperationResponse, error)
}

type ListInstanceProcessesCompleteResult struct {
	Items []ProcessInfo
}

func (r ListInstanceProcessesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessesOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcesses ...
func (c WebAppsClient) ListInstanceProcesses(ctx context.Context, id InstanceId) (resp ListInstanceProcessesOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcesses(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcesses(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcesses prepares the ListInstanceProcesses request.
func (c WebAppsClient) preparerForListInstanceProcesses(ctx context.Context, id InstanceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/processes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListInstanceProcessesWithNextLink prepares the ListInstanceProcesses request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcesses handles the response to the ListInstanceProcesses request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcesses(resp *http.Response) (result ListInstanceProcessesOperationResponse, err error) {
	type page struct {
		Values   []ProcessInfo `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessesOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcesses(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcesses", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessesComplete(ctx context.Context, id InstanceId) (ListInstanceProcessesCompleteResult, error) {
	return c.ListInstanceProcessesCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// ListInstanceProcessesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessesCompleteMatchingPredicate(ctx context.Context, id InstanceId, predicate ProcessInfoOperationPredicate) (resp ListInstanceProcessesCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	page, err := c.ListInstanceProcesses(ctx, id)
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

	out := ListInstanceProcessesCompleteResult{
		Items: items,
	}
	return out, nil
}
