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

type ListProcessesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProcessesSlotOperationResponse, error)
}

type ListProcessesSlotCompleteResult struct {
	Items []ProcessInfo
}

func (r ListProcessesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProcessesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListProcessesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProcessesSlot ...
func (c WebAppsClient) ListProcessesSlot(ctx context.Context, id SlotId) (resp ListProcessesSlotOperationResponse, err error) {
	req, err := c.preparerForListProcessesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProcessesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProcessesSlot prepares the ListProcessesSlot request.
func (c WebAppsClient) preparerForListProcessesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// preparerForListProcessesSlotWithNextLink prepares the ListProcessesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListProcessesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProcessesSlot handles the response to the ListProcessesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProcessesSlot(resp *http.Response) (result ListProcessesSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProcessesSlotOperationResponse, err error) {
			req, err := c.preparerForListProcessesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProcessesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProcessesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProcessesSlotComplete(ctx context.Context, id SlotId) (ListProcessesSlotCompleteResult, error) {
	return c.ListProcessesSlotCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// ListProcessesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProcessesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ProcessInfoOperationPredicate) (resp ListProcessesSlotCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	page, err := c.ListProcessesSlot(ctx, id)
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

	out := ListProcessesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
