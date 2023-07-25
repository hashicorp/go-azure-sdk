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

type ListSiteAnalysesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AnalysisDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteAnalysesSlotOperationResponse, error)
}

type ListSiteAnalysesSlotCompleteResult struct {
	Items []AnalysisDefinition
}

func (r ListSiteAnalysesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteAnalysesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListSiteAnalysesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteAnalysesSlot ...
func (c DiagnosticsClient) ListSiteAnalysesSlot(ctx context.Context, id SlotDiagnosticId) (resp ListSiteAnalysesSlotOperationResponse, err error) {
	req, err := c.preparerForListSiteAnalysesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteAnalysesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteAnalysesSlot prepares the ListSiteAnalysesSlot request.
func (c DiagnosticsClient) preparerForListSiteAnalysesSlot(ctx context.Context, id SlotDiagnosticId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/analyses", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteAnalysesSlotWithNextLink prepares the ListSiteAnalysesSlot request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteAnalysesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteAnalysesSlot handles the response to the ListSiteAnalysesSlot request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteAnalysesSlot(resp *http.Response) (result ListSiteAnalysesSlotOperationResponse, err error) {
	type page struct {
		Values   []AnalysisDefinition `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteAnalysesSlotOperationResponse, err error) {
			req, err := c.preparerForListSiteAnalysesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteAnalysesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalysesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteAnalysesSlotComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteAnalysesSlotComplete(ctx context.Context, id SlotDiagnosticId) (ListSiteAnalysesSlotCompleteResult, error) {
	return c.ListSiteAnalysesSlotCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// ListSiteAnalysesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteAnalysesSlotCompleteMatchingPredicate(ctx context.Context, id SlotDiagnosticId, predicate AnalysisDefinitionOperationPredicate) (resp ListSiteAnalysesSlotCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	page, err := c.ListSiteAnalysesSlot(ctx, id)
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

	out := ListSiteAnalysesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
