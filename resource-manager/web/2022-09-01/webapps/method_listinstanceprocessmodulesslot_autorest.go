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

type ListInstanceProcessModulesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProcessModuleInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceProcessModulesSlotOperationResponse, error)
}

type ListInstanceProcessModulesSlotCompleteResult struct {
	Items []ProcessModuleInfo
}

func (r ListInstanceProcessModulesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceProcessModulesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceProcessModulesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceProcessModulesSlot ...
func (c WebAppsClient) ListInstanceProcessModulesSlot(ctx context.Context, id SlotInstanceProcessId) (resp ListInstanceProcessModulesSlotOperationResponse, err error) {
	req, err := c.preparerForListInstanceProcessModulesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceProcessModulesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceProcessModulesSlot prepares the ListInstanceProcessModulesSlot request.
func (c WebAppsClient) preparerForListInstanceProcessModulesSlot(ctx context.Context, id SlotInstanceProcessId) (*http.Request, error) {
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

// preparerForListInstanceProcessModulesSlotWithNextLink prepares the ListInstanceProcessModulesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceProcessModulesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceProcessModulesSlot handles the response to the ListInstanceProcessModulesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceProcessModulesSlot(resp *http.Response) (result ListInstanceProcessModulesSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceProcessModulesSlotOperationResponse, err error) {
			req, err := c.preparerForListInstanceProcessModulesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceProcessModulesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceProcessModulesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceProcessModulesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceProcessModulesSlotComplete(ctx context.Context, id SlotInstanceProcessId) (ListInstanceProcessModulesSlotCompleteResult, error) {
	return c.ListInstanceProcessModulesSlotCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListInstanceProcessModulesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceProcessModulesSlotCompleteMatchingPredicate(ctx context.Context, id SlotInstanceProcessId, predicate ProcessModuleInfoOperationPredicate) (resp ListInstanceProcessModulesSlotCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	page, err := c.ListInstanceProcessModulesSlot(ctx, id)
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

	out := ListInstanceProcessModulesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
