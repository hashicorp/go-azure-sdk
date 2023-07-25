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

type ListInstanceProcessesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessesSlotOperationResponse, error)
}

type ListInstanceProcessesSlotCompleteResult struct {
	Items []ProcessInfo
}

func (r ListInstanceProcessesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcessesSlot ...
func (c WebAppsClient) ListInstanceProcessesSlot(ctx context.Context, id SlotInstanceId) (resp ListInstanceProcessesSlotOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcessesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcessesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcessesSlot prepares the ListInstanceProcessesSlot request.
func (c WebAppsClient) preparerForListInstanceProcessesSlot(ctx context.Context, id SlotInstanceId) (*http.Request, error) {
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

// preparerForListInstanceProcessesSlotWithNextLink prepares the ListInstanceProcessesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcessesSlot handles the response to the ListInstanceProcessesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcessesSlot(resp *http.Response) (result ListInstanceProcessesSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessesSlotOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcessesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessesSlotComplete(ctx context.Context, id SlotInstanceId) (ListInstanceProcessesSlotCompleteResult, error) {
	return c.ListInstanceProcessesSlotCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// ListInstanceProcessesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessesSlotCompleteMatchingPredicate(ctx context.Context, id SlotInstanceId, predicate ProcessInfoOperationPredicate) (resp ListInstanceProcessesSlotCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	page, err := c.ListInstanceProcessesSlot(ctx, id)
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

	out := ListInstanceProcessesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
