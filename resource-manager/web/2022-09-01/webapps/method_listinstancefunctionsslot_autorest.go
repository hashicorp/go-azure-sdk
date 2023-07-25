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

type ListInstanceFunctionsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FunctionEnvelope

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListInstanceFunctionsSlotOperationResponse, error)
}

type ListInstanceFunctionsSlotCompleteResult struct {
	Items []FunctionEnvelope
}

func (r ListInstanceFunctionsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListInstanceFunctionsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListInstanceFunctionsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListInstanceFunctionsSlot ...
func (c WebAppsClient) ListInstanceFunctionsSlot(ctx context.Context, id SlotId) (resp ListInstanceFunctionsSlotOperationResponse, err error) {
	req, err := c.preparerForListInstanceFunctionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListInstanceFunctionsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListInstanceFunctionsSlot prepares the ListInstanceFunctionsSlot request.
func (c WebAppsClient) preparerForListInstanceFunctionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/functions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListInstanceFunctionsSlotWithNextLink prepares the ListInstanceFunctionsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListInstanceFunctionsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListInstanceFunctionsSlot handles the response to the ListInstanceFunctionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListInstanceFunctionsSlot(resp *http.Response) (result ListInstanceFunctionsSlotOperationResponse, err error) {
	type page struct {
		Values   []FunctionEnvelope `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListInstanceFunctionsSlotOperationResponse, err error) {
			req, err := c.preparerForListInstanceFunctionsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListInstanceFunctionsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListInstanceFunctionsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListInstanceFunctionsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListInstanceFunctionsSlotComplete(ctx context.Context, id SlotId) (ListInstanceFunctionsSlotCompleteResult, error) {
	return c.ListInstanceFunctionsSlotCompleteMatchingPredicate(ctx, id, FunctionEnvelopeOperationPredicate{})
}

// ListInstanceFunctionsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListInstanceFunctionsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate FunctionEnvelopeOperationPredicate) (resp ListInstanceFunctionsSlotCompleteResult, err error) {
	items := make([]FunctionEnvelope, 0)

	page, err := c.ListInstanceFunctionsSlot(ctx, id)
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

	out := ListInstanceFunctionsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
