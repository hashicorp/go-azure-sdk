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

type ListProcessModulesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessModuleInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProcessModulesSlotOperationResponse, error)
}

type ListProcessModulesSlotCompleteResult struct {
	Items []ProcessModuleInfo
}

func (r ListProcessModulesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProcessModulesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListProcessModulesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProcessModulesSlot ...
func (c WebAppsClient) ListProcessModulesSlot(ctx context.Context, id SlotProcessId) (resp ListProcessModulesSlotOperationResponse, err error) {
	req, err := c.preparerForListProcessModulesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProcessModulesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProcessModulesSlot prepares the ListProcessModulesSlot request.
func (c WebAppsClient) preparerForListProcessModulesSlot(ctx context.Context, id SlotProcessId) (*http.Request, error) {
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

// preparerForListProcessModulesSlotWithNextLink prepares the ListProcessModulesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListProcessModulesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProcessModulesSlot handles the response to the ListProcessModulesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProcessModulesSlot(resp *http.Response) (result ListProcessModulesSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProcessModulesSlotOperationResponse, err error) {
			req, err := c.preparerForListProcessModulesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProcessModulesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProcessModulesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProcessModulesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProcessModulesSlotComplete(ctx context.Context, id SlotProcessId) (ListProcessModulesSlotCompleteResult, error) {
	return c.ListProcessModulesSlotCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListProcessModulesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProcessModulesSlotCompleteMatchingPredicate(ctx context.Context, id SlotProcessId, predicate ProcessModuleInfoOperationPredicate) (resp ListProcessModulesSlotCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	page, err := c.ListProcessModulesSlot(ctx, id)
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

	out := ListProcessModulesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
