package diagnostics

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

type ListSiteDetectorResponsesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DetectorResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteDetectorResponsesSlotOperationResponse, error)
}

type ListSiteDetectorResponsesSlotCompleteResult struct {
	Items []DetectorResponse
}

func (r ListSiteDetectorResponsesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteDetectorResponsesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListSiteDetectorResponsesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteDetectorResponsesSlot ...
func (c DiagnosticsClient) ListSiteDetectorResponsesSlot(ctx context.Context, id SlotId) (resp ListSiteDetectorResponsesSlotOperationResponse, err error) {
	req, err := c.preparerForListSiteDetectorResponsesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteDetectorResponsesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteDetectorResponsesSlot prepares the ListSiteDetectorResponsesSlot request.
func (c DiagnosticsClient) preparerForListSiteDetectorResponsesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectors", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteDetectorResponsesSlotWithNextLink prepares the ListSiteDetectorResponsesSlot request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteDetectorResponsesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteDetectorResponsesSlot handles the response to the ListSiteDetectorResponsesSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteDetectorResponsesSlot(resp *http.Response) (result ListSiteDetectorResponsesSlotOperationResponse, err error) {
	type page struct {
		Values   []DetectorResponse `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteDetectorResponsesSlotOperationResponse, err error) {
			req, err := c.preparerForListSiteDetectorResponsesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteDetectorResponsesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponsesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteDetectorResponsesSlotComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteDetectorResponsesSlotComplete(ctx context.Context, id SlotId) (ListSiteDetectorResponsesSlotCompleteResult, error) {
	return c.ListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListSiteDetectorResponsesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DetectorResponseOperationPredicate) (resp ListSiteDetectorResponsesSlotCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	page, err := c.ListSiteDetectorResponsesSlot(ctx, id)
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

	out := ListSiteDetectorResponsesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
