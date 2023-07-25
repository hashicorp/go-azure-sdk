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

type ListInstanceIdentifiersSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WebSiteInstanceStatus

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceIdentifiersSlotOperationResponse, error)
}

type ListInstanceIdentifiersSlotCompleteResult struct {
	Items []WebSiteInstanceStatus
}

func (r ListInstanceIdentifiersSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceIdentifiersSlotOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceIdentifiersSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceIdentifiersSlot ...
func (c WebAppsClient) ListInstanceIdentifiersSlot(ctx context.Context, id SlotId) (resp ListInstanceIdentifiersSlotOperationResponse, err error) {
	req, err := c.preparerForListInstanceIdentifiersSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceIdentifiersSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceIdentifiersSlot prepares the ListInstanceIdentifiersSlot request.
func (c WebAppsClient) preparerForListInstanceIdentifiersSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/instances", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListInstanceIdentifiersSlotWithNextLink prepares the ListInstanceIdentifiersSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceIdentifiersSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceIdentifiersSlot handles the response to the ListInstanceIdentifiersSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceIdentifiersSlot(resp *http.Response) (result ListInstanceIdentifiersSlotOperationResponse, err error) {
	type page struct {
		Values   []WebSiteInstanceStatus `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceIdentifiersSlotOperationResponse, err error) {
			req, err := c.preparerForListInstanceIdentifiersSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceIdentifiersSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceIdentifiersSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceIdentifiersSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceIdentifiersSlotComplete(ctx context.Context, id SlotId) (ListInstanceIdentifiersSlotCompleteResult, error) {
	return c.ListInstanceIdentifiersSlotCompleteMatchingPredicate(ctx, id, WebSiteInstanceStatusOperationPredicate{})
}

// ListInstanceIdentifiersSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceIdentifiersSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WebSiteInstanceStatusOperationPredicate) (resp ListInstanceIdentifiersSlotCompleteResult, err error) {
	items := make([]WebSiteInstanceStatus, 0)

	page, err := c.ListInstanceIdentifiersSlot(ctx, id)
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

	out := ListInstanceIdentifiersSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
